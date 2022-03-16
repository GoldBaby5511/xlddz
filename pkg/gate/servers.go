package gate

import (
	"github.com/golang/protobuf/proto"
	"mango/api/center"
	"mango/pkg/conf"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
	"reflect"
	"time"
)

type agentServer struct {
	tcpClient *n.TCPClient
	conn      n.Conn
	info      n.BaseAgentInfo
}

func newServerItem(info n.BaseAgentInfo, autoReconnect bool, pendingWriteNum int) {
	if info.ListenOnAddr == "" {
		log.Warning("agentServer", "警告,没地址怎么连接?,info=%v,autoReconnect=%v,pendingWriteNum=%v",
			info, autoReconnect, pendingWriteNum)
		return
	}

	tcpClient := new(n.TCPClient)
	tcpClient.Addr = info.ListenOnAddr
	tcpClient.PendingWriteNum = pendingWriteNum
	tcpClient.AutoReconnect = autoReconnect
	tcpClient.NewAgent = func(conn *n.TCPConn) n.AgentServer {
		tcpClient.AutoReconnect = true
		a := &agentServer{tcpClient: tcpClient, conn: conn, info: info}
		log.Debug("agentServer", "连接成功,info=%v", a.info)
		sendRegAppReq(a)
		timeInterval := 30 * time.Second
		timerHeartbeat := time.NewTimer(timeInterval)
		go func(t *time.Timer) {
			for {
				<-t.C
				var pulse center.AppPulseNotify
				pulse.Action = (*center.AppPulseNotify_PulseAction)(proto.Int32(int32(center.AppPulseNotify_HeartBeatReq)))
				pulse.PulseData = proto.Uint64(uint64(time.Now().Unix()))
				a.SendData(n.AppCenter, uint32(center.CMDCenter_IDPulseNotify), &pulse)

				t.Reset(timeInterval)
			}
		}(timerHeartbeat)

		mxServers.Lock()
		servers[util.MakeUint64FromUint32(info.AppType, info.AppID)] = a
		mxServers.Unlock()

		sendRegConfigReq()
		return a
	}

	log.Debug("agentServer", "开始连接,info=%v", info)

	tcpClient.Start()
}

func (a *agentServer) Run() {
	for {
		bm, msgData, err := a.conn.ReadMsg()
		if err != nil {
			log.Warning("agentServer", "异常,网关读取消息失败,info=%v,err=%v", a.info, err)
			break
		}

		if bm.Cmd.AppType != uint16(n.AppCenter) {
			log.Warning("agentServer", "不可能出现非center消息,cmd=%v", bm.Cmd)
			break
		}

		switch bm.Cmd.CmdId {
		case uint16(center.CMDCenter_IDAppRegRsp):
			var m center.RegisterAppRsp
			_ = proto.Unmarshal(msgData, &m)

			log.Info("agentServer", "注册消息,regResult=%v,CenterId=%v,appName=%v,appType=%v,appId=%v,Addr=%v",
				m.GetRegResult(), m.GetCenterId(), m.GetAppName(), m.GetAppType(), m.GetAppId(), m.GetAppAddress())

			if m.GetRegResult() == 0 {
				mxServers.Lock()
				_, ok := servers[util.MakeUint64FromUint32(m.GetAppType(), m.GetAppId())]
				mxServers.Unlock()

				if !(conf.AppInfo.Type == m.GetAppType() && conf.AppInfo.Id == m.GetAppId()) && m.GetAppAddress() != "" && !ok {
					info := n.BaseAgentInfo{AgentType: n.CommonServer, AppName: m.GetAppName(), AppType: m.GetAppType(), AppID: m.GetAppId(), ListenOnAddr: m.GetAppAddress()}
					newServerItem(info, false, 0)
				}

				if conf.AppInfo.Type == n.AppConfig {
					mxServers.Lock()
					key := util.MakeUint64FromUint32(n.AppCenter, 0)
					if _, ok := servers[key]; ok {
						servers[key].info.AppID = m.GetCenterId()
					}
					mxServers.Unlock()
				}
			}

			if agentChanRPC != nil {
				agentChanRPC.Call0(CenterRegResult, m.GetRegResult(), m.GetCenterId())
			}
		case uint16(center.CMDCenter_IDAppState): //app状态改变
			var m center.AppStateNotify
			_ = proto.Unmarshal(msgData, &m)
			log.Debug("agentServer", "app状态改变 AppState=%v,CenterId=%v,AppType=%v,AppId=%v",
				m.GetAppState(), m.GetCenterId(), m.GetAppType(), m.GetAppId())

			mxServers.Lock()
			key := util.MakeUint64FromUint32(m.GetAppType(), m.GetAppId())
			if _, ok := servers[key]; ok {
				servers[key].Close()
			}
			mxServers.Unlock()

		case uint16(center.CMDCenter_IDPulseNotify):
		default:
			log.Error("agentServer", "n.AppCenter,异常,还未处理消息,%v", bm.Cmd)
		}
	}
}

func (a *agentServer) OnClose() {
	log.Debug("", "服务间连接断开了,info=%v", a.info)
	if a.info.AppType == n.AppLogger {
		log.SetCallback(nil)
		log.Info("agentServer", "日志服务器断开")
	} else if a.info.AppType == n.AppCenter {
		log.Warning("agentServer", "异常,与center连接断开,世界需要重启... ...")
		for _, c := range cbCenterDisconnect {
			c()
		}
	}
	if a.tcpClient != nil && !a.tcpClient.AutoReconnect {
		a.tcpClient.Close()
	}

	if !a.tcpClient.AutoReconnect {
		mxServers.Lock()
		delete(servers, util.MakeUint64FromUint32(a.info.AppType, a.info.AppID))
		mxServers.Unlock()
	}
}

func (a *agentServer) SendMessage(bm n.BaseMessage) {
	m := bm.MyMessage.(proto.Message)
	data, err := proto.Marshal(m)
	if err != nil {
		log.Error("agentServer", "异常,proto.Marshal %v error: %v", reflect.TypeOf(m), err)
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
		log.Error("agentServer", "写信息失败 %v error: %v", reflect.TypeOf(m), err)
	}
}

func (a *agentServer) SendData(appType, cmdId uint32, m proto.Message) {
	data, err := proto.Marshal(m)
	if err != nil {
		log.Error("agentServer", "异常,proto.Marshal %v error: %v", reflect.TypeOf(m), err)
		return
	}
	err = a.conn.WriteMsg(uint16(appType), uint16(cmdId), data, nil)
	if err != nil {
		log.Error("agentServer", "write message %v error: %v", reflect.TypeOf(m), err)
	}
}

func (a *agentServer) Close() {
	a.conn.Close()
}
func (a *agentServer) Destroy() {
	a.conn.Destroy()
}
