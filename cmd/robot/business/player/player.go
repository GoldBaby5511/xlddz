package player

import (
	"github.com/golang/protobuf/proto"
	"mango/api/client"
	"mango/api/gate"
	"mango/api/types"
	"mango/pkg/conf"
	"mango/pkg/log"
	"mango/pkg/module"
	n "mango/pkg/network"
	"mango/pkg/network/protobuf"
	"mango/pkg/util"
	"reflect"
)

const (
	NilState       uint32 = 0
	Logging        uint32 = 1
	LoggedIn       uint32 = 2
	JoinRoom       uint32 = 3
	StandingInRoom uint32 = 4
	HandsUp        uint32 = 5
	PlayingState   uint32 = 6
)

type Player struct {
	a         *agentPlayer
	processor *protobuf.Processor
	Skeleton  *module.Skeleton
	roomList  map[uint64]*types.RoomInfo
	Account   string
	PassWord  string
	UserID    uint64
	State     uint32
	RoomID    uint32
}

func NewPlayer() *Player {
	p := new(Player)
	p.UserID = 0
	p.State = NilState
	p.roomList = make(map[uint64]*types.RoomInfo)
	p.processor = protobuf.NewProcessor()
	p.Skeleton = module.NewSkeleton(conf.GoLen, conf.TimerDispatcherLen, conf.AsynCallLen, conf.ChanRPCLen)
	go func() {
		p.Skeleton.Run()
	}()

	p.MsgRegister(&gate.HelloRsp{}, n.CMDGate, uint16(gate.CMDID_Gate_IDHelloRsp), p.handleHelloRsp)
	p.MsgRegister(&client.LoginRsp{}, n.CMDClient, uint16(client.CMDID_Client_IDLoginRsp), p.handleLoginRsp)
	p.MsgRegister(&client.RoomListRsp{}, n.CMDClient, uint16(client.CMDID_Client_IDRoomListRsp), p.handleRoomListRsp)
	p.MsgRegister(&client.JoinRoomRsp{}, n.CMDClient, uint16(client.CMDID_Client_IDJoinRoomRsp), p.handleJoinRoomRsp)
	p.MsgRegister(&client.RoomActionRsp{}, n.CMDClient, uint16(client.CMDID_Client_IDRoomActionRsp), p.handleRoomActionRsp)
	return p
}

func (p *Player) MsgRegister(m proto.Message, mainCmdId uint32, subCmdId uint16, f interface{}) {
	chanRPC := p.Skeleton.ChanRPCServer
	p.processor.Register(m, mainCmdId, subCmdId, chanRPC)
	chanRPC.Register(reflect.TypeOf(m), f)
}

func (p *Player) Connect() {
	tcpClient := new(n.TCPClient)
	tcpClient.Addr = "127.0.0.1:10102"
	tcpClient.PendingWriteNum = 0
	tcpClient.AutoReconnect = false
	tcpClient.NewAgent = func(conn *n.TCPConn) n.AgentServer {
		a := &agentPlayer{tcpClient: tcpClient, conn: conn, p: p}
		log.Debug("agentServer", "连接成功,%v", a.info)
		p.connectSuccess(a)
		return a
	}

	log.Debug("agentServer", "开始连接")

	if tcpClient != nil {
		tcpClient.Start()
	}
}

func (p *Player) CheckRoomList() {
	var req client.RoomListReq
	req.ListId = proto.Uint32(0)
	cmd := n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDRoomListReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd}
	p.SendMessage2Gate(n.AppList, n.Send2AnyOne, bm)
}

func (p *Player) JoinRoom() {
	if len(p.roomList) == 0 {
		return
	}
	r := &types.RoomInfo{}
	for _, v := range p.roomList {
		r = v
		break
	}

	log.Debug("", "进入房间,a=%v,p=%v", p.Account, p.PassWord)

	p.State = JoinRoom
	var req client.JoinRoomReq
	cmd := n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDJoinRoomReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd}
	p.SendMessage2Gate(r.RoomInfo.GetType(), r.RoomInfo.GetId(), bm)
}

func (p *Player) ActionRoom() {
	if p.State != StandingInRoom {
		return
	}

	log.Debug("", "房间动作,a=%v,p=%v", p.Account, p.PassWord)

	var req client.RoomActionReq
	req.Action = (*client.ActionType)(proto.Int32(int32(client.ActionType_Ready)))
	cmd := n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDRoomActionReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd}
	p.SendMessage2Gate(n.AppRoom, p.RoomID, bm)
}

func (p *Player) handleHelloRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*gate.HelloRsp)
	//a := args[n.AgentIndex].(n.AgentClient)

	log.Debug("hello", "收到hello消息,RspFlag=%v", m.GetRspFlag())

	var req client.LoginReq
	req.LoginAccount = proto.String(p.Account)
	req.LoginPassword = proto.String(p.PassWord)
	cmd := n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDLoginReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd, TraceId: b.TraceId}
	p.SendMessage2Gate(n.AppLogin, n.Send2AnyOne, bm)
}

func (p *Player) handleLoginRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.LoginRsp)
	//a := args[n.AgentIndex].(n.AgentClient)

	log.Debug("hello", "收到登录回复消息,Result=%v", m.GetLoginResult())
	if m.GetLoginResult() == client.LoginRsp_SUCCESS {
		p.State = LoggedIn
	}
}

func (p *Player) handleRoomListRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.RoomListRsp)
	//a := args[n.AgentIndex].(n.AgentClient)

	log.Debug("hello", "收到列表消息,len=%v", len(m.GetRooms()))
	for _, r := range m.GetRooms() {
		regKey := util.MakeUint64FromUint32(r.GetRoomInfo().GetType(), r.GetRoomInfo().GetId())
		p.roomList[regKey] = r
	}
}

func (p *Player) handleJoinRoomRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.JoinRoomRsp)
	//a := args[n.AgentIndex].(n.AgentClient)

	log.Debug("hello", "收到加入消息,Code=%v", m.GetErrInfo().GetCode())
	if m.GetErrInfo().GetCode() == 0 {
		p.State = StandingInRoom
		p.RoomID = m.GetAppId()
	}
}

func (p *Player) handleRoomActionRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.RoomActionRsp)
	//a := args[n.AgentIndex].(n.AgentClient)

	log.Debug("hello", "收到动作消息消息,Code=%v", m.GetErrInfo().GetCode())
	if m.GetErrInfo().GetCode() == 0 {
		p.State = HandsUp
	}
}

func (p *Player) connectSuccess(a *agentPlayer) {
	p.a = a

	var req gate.HelloReq
	p.a.SendData(n.CMDGate, uint32(gate.CMDID_Gate_IDHelloReq), &req)

	//var dataReq gate.TransferDataReq
	//dataReq.AttApptype = proto.Uint32(n.AppGate)
	//dataReq.AttAppid = proto.Uint32(uint32(gateConnID >> 32))
	//dataReq.DataCmdKind = proto.Uint32(uint32(bm.Cmd.MainCmdID))
	//dataReq.DataCmdSubid = proto.Uint32(uint32(bm.Cmd.SubCmdID))
	//dataReq.Data, _ = proto.Marshal(&req)
	//dataReq.Gateconnid = proto.Uint64(gateConnID)
	//dataReq.AttSessionid = proto.Uint64(sessionID)
	//cmd := n.TCPCommand{MainCmdID: uint16(n.CMDGate), SubCmdID: uint16(gate.CMDID_Gate_IDTransferDataReq)}
	//transBM := n.BaseMessage{MyMessage: &dataReq, Cmd: cmd, TraceId: bm.TraceId}
}

func (p *Player) SendMessage2Gate(destAppType, destAppid uint32, bm n.BaseMessage) {
	var dataReq gate.TransferDataReq
	dataReq.AttApptype = proto.Uint32(destAppType)
	dataReq.AttAppid = proto.Uint32(destAppid)
	dataReq.DataCmdKind = proto.Uint32(uint32(bm.Cmd.MainCmdID))
	dataReq.DataCmdSubid = proto.Uint32(uint32(bm.Cmd.SubCmdID))
	dataReq.Data, _ = proto.Marshal(bm.MyMessage.(proto.Message))
	dataReq.Gateconnid = proto.Uint64(0)
	cmd := n.TCPCommand{MainCmdID: uint16(n.CMDGate), SubCmdID: uint16(gate.CMDID_Gate_IDTransferDataReq)}
	transBM := n.BaseMessage{MyMessage: &dataReq, Cmd: cmd, TraceId: bm.TraceId}
	p.a.SendMessage(transBM)
	//g.SendData2App(n.AppGate, n.Send2AnyOne, n.CMDGate, uint32(gate.CMDID_Gate_IDTransferDataReq), &dataReq)
}

type agentPlayer struct {
	tcpClient *n.TCPClient
	p         *Player
	conn      n.Conn
	info      n.BaseAgentInfo
}

func (a *agentPlayer) Run() {
	for {
		bm, msgData, err := a.conn.ReadMsg()
		if err != nil {
			log.Warning("agentPlayer", "异常,网关读取消息失败,info=%v,err=%v", a.info, err)
			break
		}

		if a.p.processor == nil {
			log.Warning("", "processor==nil,cmd=%v", bm.Cmd)
			break
		}

		log.Debug("", "收到消息,bm=%v,len=%v", bm, len(msgData))

		unmarshalCmd := bm.Cmd
		var cmd, msg, dataReq interface{}
		if bm.Cmd.MainCmdID == uint16(n.CMDGate) && bm.Cmd.SubCmdID == uint16(gate.CMDID_Gate_IDTransferDataReq) && conf.AppInfo.AppType != n.AppGate {
			var m gate.TransferDataReq
			_ = proto.Unmarshal(msgData, &m)
			unmarshalCmd = n.TCPCommand{MainCmdID: uint16(m.GetDataCmdKind()), SubCmdID: uint16(m.GetDataCmdSubid())}
			msgData = m.GetData()
			dataReq = &m
		} else {
			dataReq = a.info
		}

		cmd, msg, err = a.p.processor.Unmarshal(unmarshalCmd.MainCmdID, unmarshalCmd.SubCmdID, msgData)
		if err != nil {
			log.Error("agentClient", "unmarshal message,headCmd=%v,error: %v", bm.Cmd, err)
			continue
		}
		err = a.p.processor.Route(n.BaseMessage{MyMessage: msg, TraceId: bm.TraceId}, a, cmd, dataReq)
		if err != nil {
			log.Error("agentClient", "client agentClient route message error: %v,cmd=%v", err, cmd)
			continue
		}
	}
}

func (a *agentPlayer) OnClose() {
	log.Debug("", "服务间连接断开了,info=%v", a.info)
}

func (a *agentPlayer) SendMessage(bm n.BaseMessage) {
	m := bm.MyMessage.(proto.Message)
	data, err := proto.Marshal(m)
	if err != nil {
		log.Error("agentPlayer", "异常,proto.Marshal %v error: %v", reflect.TypeOf(m), err)
		return
	}
	//追加TraceId
	otherData := make([]byte, 0, n.TraceIdLen+1)
	if bm.TraceId != "" {
		otherData = append(otherData, n.FlagOtherTraceId)
		otherData = append(otherData, []byte(bm.TraceId)...)
	}
	err = a.conn.WriteMsg(bm.Cmd.MainCmdID, bm.Cmd.SubCmdID, data, otherData)
	if err != nil {
		log.Error("agentPlayer", "写信息失败 %v error: %v", reflect.TypeOf(m), err)
	}
}

func (a *agentPlayer) SendData(mainCmdID, subCmdID uint32, m proto.Message) {
	data, err := proto.Marshal(m)
	if err != nil {
		log.Error("agentPlayer", "异常,proto.Marshal %v error: %v", reflect.TypeOf(m), err)
		return
	}
	err = a.conn.WriteMsg(uint16(mainCmdID), uint16(subCmdID), data, nil)
	if err != nil {
		log.Error("agentPlayer", "write message %v error: %v", reflect.TypeOf(m), err)
	}
}

func (a *agentPlayer) Close() {
	a.conn.Close()
}
func (a *agentPlayer) Destroy() {
	a.conn.Destroy()
}
