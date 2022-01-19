package player

import (
	"github.com/golang/protobuf/proto"
	"mango/api/gameddz"
	"mango/api/gateway"
	"mango/api/list"
	"mango/api/login"
	"mango/api/room"
	tCMD "mango/api/table"
	"mango/api/types"
	"mango/pkg/conf"
	"mango/pkg/conf/apollo"
	"mango/pkg/log"
	"mango/pkg/module"
	n "mango/pkg/network"
	"mango/pkg/network/protobuf"
	"mango/pkg/timer"
	"mango/pkg/util"
	"math/rand"
	"reflect"
	"strconv"
	"time"
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

const (
	GameBegin uint32 = 1
	GameOver  uint32 = 2
)

type Player struct {
	a              *agentPlayer
	processor      *protobuf.Processor
	Skeleton       *module.Skeleton
	roomList       map[uint64]*types.RoomInfo
	Account        string
	PassWord       string
	UserId         uint64
	State          uint32
	RoomID         uint32
	TableServiceId uint32
	TableId        uint64
	SeatId         uint32
	Scene          uint32
}

func NewPlayer(account, passWord string) *Player {
	p := new(Player)
	p.Account = account
	p.PassWord = passWord
	p.UserId = 0
	p.State = NilState
	p.roomList = make(map[uint64]*types.RoomInfo)
	p.processor = protobuf.NewProcessor()
	p.Skeleton = module.NewSkeleton(conf.GoLen, conf.TimerDispatcherLen, conf.AsynCallLen, conf.ChanRPCLen)
	go func() {
		p.Skeleton.Run()
	}()

	p.MsgRegister(&gateway.HelloRsp{}, n.AppGate, uint16(gateway.CMDGateway_IDHelloRsp), p.handleHelloRsp)
	p.MsgRegister(&login.LoginRsp{}, n.AppLogin, uint16(login.CMDLogin_IDLoginRsp), p.handleLoginRsp)
	p.MsgRegister(&list.RoomListRsp{}, n.AppList, uint16(list.CMDList_IDRoomListRsp), p.handleRoomListRsp)
	p.MsgRegister(&room.JoinRsp{}, n.AppRoom, uint16(room.CMDRoom_IDJoinRsp), p.handleJoinRoomRsp)
	p.MsgRegister(&room.UserActionRsp{}, n.AppRoom, uint16(room.CMDRoom_IDUserActionRsp), p.handleRoomActionRsp)
	p.MsgRegister(&room.UserStateChange{}, n.AppRoom, uint16(room.CMDRoom_IDUserStateChange), p.handleUserStateChange)

	p.MsgRegister(&gameddz.GameStart{}, n.AppTable, uint16(gameddz.CMDGameddz_IDGameStart), p.handleGameStart)
	p.MsgRegister(&gameddz.OutCardRsp{}, n.AppTable, uint16(gameddz.CMDGameddz_IDOutCardRsp), p.handleOutCardRsp)
	p.MsgRegister(&gameddz.GameOver{}, n.AppTable, uint16(gameddz.CMDGameddz_IDGameOver), p.handleGameOver)

	p.Skeleton.AfterFunc(time.Duration(rand.Intn(9)+1)*time.Second, p.connect)

	return p
}

func (p *Player) MsgRegister(m proto.Message, mainCmdId uint32, subCmdId uint16, f interface{}) {
	chanRPC := p.Skeleton.ChanRPCServer
	p.processor.Register(m, mainCmdId, subCmdId, chanRPC)
	chanRPC.Register(reflect.TypeOf(m), f)
}

func (p *Player) heartbeat() {
	var req gateway.PulseReq
	p.a.SendData(n.AppGate, uint32(gateway.CMDGateway_IDPulseReq), &req)
}

func (p *Player) connect() {
	tcpClient := new(n.TCPClient)
	tcpClient.Addr = apollo.GetConfig("网关地址", "127.0.0.1:10100")
	if conf.RunInLocalDocker() {
		tcpClient.Addr = "gateway:" + strconv.Itoa(util.GetPortFromIPAddress(tcpClient.Addr))
	}
	tcpClient.PendingWriteNum = 0
	tcpClient.AutoReconnect = false
	tcpClient.NewAgent = func(conn *n.TCPConn) n.AgentServer {
		a := &agentPlayer{tcpClient: tcpClient, conn: conn, p: p}
		p.a = a

		var req gateway.HelloReq
		a.SendData(n.AppGate, uint32(gateway.CMDGateway_IDHelloReq), &req)
		return a
	}

	log.Debug("", "开始连接,UserId=%v,a=%v", p.UserId, p.Account)

	if tcpClient != nil {
		tcpClient.Start()
	}
}

func (p *Player) checkRoomList() {
	var req list.RoomListReq
	req.ListId = proto.Uint32(0)
	cmd := n.TCPCommand{AppType: uint16(n.AppList), CmdId: uint16(list.CMDList_IDRoomListReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd}
	p.SendMessage2Gate(n.AppList, n.Send2AnyOne, bm)
}

func (p *Player) joinRoom() {
	if len(p.roomList) == 0 {
		return
	}
	r := &types.RoomInfo{}
	for _, v := range p.roomList {
		r = v
		break
	}

	log.Debug("", "进入房间,UserId=%v,a=%v,p=%v", p.UserId, p.Account, p.PassWord)

	p.State = JoinRoom
	var req room.JoinReq
	cmd := n.TCPCommand{AppType: uint16(n.AppRoom), CmdId: uint16(room.CMDRoom_IDJoinReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd}
	p.SendMessage2Gate(r.AppInfo.GetType(), r.AppInfo.GetId(), bm)
}

func (p *Player) ActionRoom() {
	if p.State != StandingInRoom {
		return
	}

	log.Debug("", "房间动作,UserId=%v,a=%v,p=%v", p.UserId, p.Account, p.PassWord)

	var req room.UserActionReq
	req.Action = (*room.ActionType)(proto.Int32(int32(room.ActionType_Ready)))
	cmd := n.TCPCommand{AppType: uint16(n.AppRoom), CmdId: uint16(room.CMDRoom_IDUserActionReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd}
	p.SendMessage2Gate(n.AppRoom, p.RoomID, bm)
}

func (p *Player) handleHelloRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*gateway.HelloRsp)

	log.Debug("", "收到hello消息,UserId=%v,a=%v,RspFlag=%v", p.UserId, p.Account, m.GetRspFlag())

	var req login.LoginReq
	req.Account = proto.String(p.Account)
	req.Password = proto.String(p.PassWord)
	cmd := n.TCPCommand{AppType: uint16(n.AppLogin), CmdId: uint16(login.CMDLogin_IDLoginReq)}
	bm := n.BaseMessage{MyMessage: &req, Cmd: cmd, TraceId: b.TraceId}
	p.SendMessage2Gate(n.AppLogin, n.Send2AnyOne, bm)

	p.Skeleton.LoopFunc(30*time.Second, p.heartbeat, timer.LoopForever)
}

func (p *Player) handleLoginRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*login.LoginRsp)

	p.UserId = m.GetBaseInfo().GetUserId()

	log.Debug("", "收到登录回复,UserId=%v,a=%v,Result=%v", p.UserId, p.Account, m.GetResult())
	if m.GetResult() == login.LoginRsp_SUCCESS {
		p.State = LoggedIn
		p.Skeleton.AfterFunc(time.Duration(rand.Intn(3)+1)*time.Second, p.checkRoomList)
	}
}

func (p *Player) handleRoomListRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*list.RoomListRsp)

	log.Debug("", "收到列表,UserId=%v,a=%v,len=%v", p.UserId, p.Account, len(m.GetRooms()))
	for _, r := range m.GetRooms() {
		regKey := util.MakeUint64FromUint32(r.AppInfo.GetType(), r.AppInfo.GetId())
		p.roomList[regKey] = r
		p.Skeleton.AfterFunc(time.Duration(rand.Intn(3)+1)*time.Second, p.joinRoom)
	}
}

func (p *Player) handleJoinRoomRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*room.JoinRsp)

	log.Debug("", "收到加入,UserId=%v,a=%v,Code=%v", p.UserId, p.Account, m.GetErrInfo().GetCode())
	if m.GetErrInfo().GetCode() == 0 {
		p.State = StandingInRoom
		p.RoomID = m.GetAppId()
		p.Skeleton.AfterFunc(time.Duration(rand.Intn(3)+1)*time.Second, p.ActionRoom)
	}
}

func (p *Player) handleRoomActionRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*room.UserActionRsp)

	log.Debug("hello", "收到动作,UserId=%v,a=%v,Code=%v", p.UserId, p.Account, m.GetErrInfo().GetCode())
	if m.GetErrInfo().GetCode() == 0 {
		p.State = HandsUp
	}
}

func (p *Player) handleUserStateChange(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*room.UserStateChange)

	log.Debug("hello", "状态变化,UserId=%v,a=%v,State=%v,tableId=%v,seatId=%v",
		p.UserId, p.Account, m.GetUserState(), m.GetTableId(), m.GetSeatId())

	p.TableServiceId = m.GetTableServiceId()
	p.TableId = m.GetTableId()
	p.SeatId = m.GetSeatId()
}

func (p *Player) handleGameStart(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*gameddz.GameStart)

	p.Scene = GameBegin
	log.Debug("", "游戏开始,UserId=%v,a=%v,TableId=%v,CurrentSeat=%v,p.SeatId=%v", p.UserId, p.Account, p.TableId, m.GetCurrentSeat(), p.SeatId)

	if p.SeatId == m.GetCurrentSeat() {
		p.Skeleton.AfterFunc(time.Duration(rand.Intn(3)+1)*time.Second, p.outCards)
	}
}

func (p *Player) handleOutCardRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*gameddz.OutCardRsp)

	log.Debug("", "收到出牌消息,UserId=%v,a=%v,CurrentSeat=%v,SeatId=%v", p.UserId, p.Account, m.GetCurrentSeat(), p.SeatId)

	if p.SeatId == m.GetCurrentSeat() {
		p.Skeleton.AfterFunc(time.Duration(rand.Intn(3)+1)*time.Second, p.outCards)
	}
}

func (p *Player) handleGameOver(args []interface{}) {
	//b := args[n.DataIndex].(n.BaseMessage)
	//m := (b.MyMessage).(*gameddz.GameOver)
	//a := args[n.AgentIndex].(n.AgentClient)
	p.Scene = GameOver
	log.Debug("", "游戏结束消息,UserId=%v,a=%v", p.UserId, p.Account)

	p.State = StandingInRoom
	p.Skeleton.AfterFunc(time.Duration(rand.Intn(3)+3)*time.Second, p.ActionRoom)
}

func (p *Player) outCards() {

	if p.Scene == GameOver {
		return
	}

	log.Debug("", "出牌,UserId=%v,a=%v,SeatId=%v,Scene=%v", p.UserId, p.Account, p.SeatId, p.Scene)

	var gameMessage tCMD.GameMessage
	gameMessage.SubCmdid = proto.Uint32(uint32(gameddz.CMDGameddz_IDOutCardReq))
	var req gameddz.OutCardReq
	for i := 0; i < rand.Intn(3)+1; i++ {
		req.OutCard = append(req.OutCard, byte(rand.Intn(3)+1))
	}
	gameMessage.Data, _ = proto.Marshal(&req)
	cmd := n.TCPCommand{AppType: uint16(n.AppTable), CmdId: uint16(tCMD.CMDTable_IDGameMessage)}
	bm := n.BaseMessage{MyMessage: &gameMessage, Cmd: cmd}
	p.SendMessage2Gate(n.AppTable, p.TableServiceId, bm)
}

func (p *Player) SendMessage2Gate(destAppType, destAppid uint32, bm n.BaseMessage) {
	var dataReq gateway.TransferDataReq
	dataReq.AttApptype = proto.Uint32(destAppType)
	dataReq.AttAppid = proto.Uint32(destAppid)
	dataReq.DataCmdKind = proto.Uint32(uint32(bm.Cmd.AppType))
	dataReq.DataCmdSubid = proto.Uint32(uint32(bm.Cmd.CmdId))
	dataReq.Data, _ = proto.Marshal(bm.MyMessage.(proto.Message))
	dataReq.Gateconnid = proto.Uint64(0)
	cmd := n.TCPCommand{AppType: uint16(n.AppGate), CmdId: uint16(gateway.CMDGateway_IDTransferDataReq)}
	transBM := n.BaseMessage{MyMessage: &dataReq, Cmd: cmd, TraceId: bm.TraceId}
	p.a.SendMessage(transBM)
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

		unmarshalCmd := bm.Cmd
		var cmd, msg, dataReq interface{}
		if bm.Cmd.AppType == uint16(n.AppGate) && bm.Cmd.CmdId == uint16(gateway.CMDGateway_IDTransferDataReq) && conf.AppInfo.Type != n.AppGate {
			var m gateway.TransferDataReq
			_ = proto.Unmarshal(msgData, &m)
			unmarshalCmd = n.TCPCommand{AppType: uint16(m.GetDataCmdKind()), CmdId: uint16(m.GetDataCmdSubid())}
			msgData = m.GetData()
			dataReq = &m
		} else {
			dataReq = a.info
		}

		cmd, msg, err = a.p.processor.Unmarshal(unmarshalCmd.AppType, unmarshalCmd.CmdId, msgData)
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
	err = a.conn.WriteMsg(bm.Cmd.AppType, bm.Cmd.CmdId, data, otherData)
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
