package gate

import (
	"github.com/golang/protobuf/proto"
	"mango/api/center"
	"mango/api/gateway"
	"mango/pkg/conf"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
	"net"
	"reflect"
)

//代理
type agentClient struct {
	id   uint32
	conn n.Conn
	info n.BaseAgentInfo
}

func (a *agentClient) Run() {
	for {
		bm, msgData, err := a.conn.ReadMsg()
		if err != nil {
			log.Warning("agentClient", "异常,网关读取消息失败,id=%v,err=%v", a.id, err)
			break
		}
		if processor == nil {
			log.Error("agentClient", "异常,解析器为nil断开连接,cmd=%v", &bm.Cmd)
			break
		}
		if conf.AppInfo.Type != n.AppCenter && bm.Cmd.AppType == uint16(n.AppCenter) {
			if bm.Cmd.CmdId == uint16(center.CMDCenter_IDAppRegReq) {
				var m center.RegisterAppReq
				_ = proto.Unmarshal(msgData, &m)
				a.info = n.BaseAgentInfo{AgentType: n.CommonServer, AppName: m.GetAppName(), AppType: m.GetAppType(), AppID: m.GetAppId()}
				if agentChanRPC != nil {
					agentChanRPC.Call0(CommonServerReg, a, a.info)
				}
				log.Debug("", "相互注册,%v", a.info)
			} else if bm.Cmd.CmdId == uint16(center.CMDCenter_IDPulseNotify) {

			}
			continue
		}

		unmarshalCmd := bm.Cmd
		var cmd, msg, dataReq interface{}
		if bm.Cmd.AppType == uint16(n.AppGate) && bm.Cmd.CmdId == uint16(gateway.CMDGateway_IDTransferDataReq) && conf.AppInfo.Type != n.AppGate {
			var m gateway.TransferDataReq
			_ = proto.Unmarshal(msgData, &m)
			unmarshalCmd = n.TCPCommand{AppType: uint16(m.GetDataApptype()), CmdId: uint16(m.GetDataCmdid())}
			msgData = m.GetData()
			dataReq = &m
			bm.AgentInfo = n.BaseAgentInfo{AgentType: n.NormalUser, AppType: util.GetHUint32FromUint64(m.GetGateconnid()), AppID: util.GetLUint32FromUint64(m.GetGateconnid())}
		} else {
			bm.AgentInfo = a.info
			dataReq = a.info
		}

		cmd, msg, err = processor.Unmarshal(unmarshalCmd.AppType, unmarshalCmd.CmdId, msgData)
		if err != nil {
			log.Error("agentClient", "unmarshal message,headCmd=%v,error: %v", bm.Cmd, err)
			continue
		}
		err = processor.Route(n.BaseMessage{MyMessage: msg, AgentInfo: bm.AgentInfo, TraceId: bm.TraceId}, a, cmd, dataReq)
		if err != nil {
			log.Error("agentClient", "client agentClient route message error: %v,cmd=%v", err, cmd)
			continue
		}
	}
}

func (a *agentClient) OnClose() {
	if agentChanRPC != nil {
		err := agentChanRPC.Call0(Disconnect, a, a.id)
		if err != nil {
			log.Warning("agentClient", "agentClient OnClose err=%v", err)
		}
	}
}

func (a *agentClient) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agentClient) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agentClient) Close() {
	a.conn.Close()
}

func (a *agentClient) Destroy() {
	a.conn.Destroy()
}

func (a *agentClient) SendData(appType, cmdId uint32, m proto.Message) {
	data, err := proto.Marshal(m)
	if err != nil {
		log.Error("agentClient", "异常,proto.Marshal %v error: %v", reflect.TypeOf(m), err)
		return
	}
	err = a.conn.WriteMsg(uint16(appType), uint16(cmdId), data, nil)
	if err != nil {
		log.Error("agentClient", "write message %v error: %v", reflect.TypeOf(m), err)
	}
}

func (a *agentClient) AgentInfo() n.BaseAgentInfo {
	return a.info
}
