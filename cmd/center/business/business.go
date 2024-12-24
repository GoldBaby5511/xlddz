package business

import (
	"encoding/json"
	"fmt"
	"mango/api/center"
	lconf "mango/pkg/conf"
	"mango/pkg/conf/apollo"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
	"math/rand"
	"time"
)

var (
	appConnData = make(map[n.AgentClient]*connectionData)
	appRegData  = make(map[uint64]*connectionData)
)

// 连接数据
type connectionData struct {
	a                  n.AgentClient
	appInfo            lconf.BaseInfo
	appState           int
	regToken           string
	stateDescription   string
	totalHearDelayTime int64
	lastHeartbeat      int64
	httpAddr           string
	rpcAddr            string
}

func init() {
	g.MsgRegister(&center.RegisterAppReq{}, n.AppCenter, uint16(center.CMDCenter_IDAppRegReq), handleRegisterAppReq)
	g.MsgRegister(&center.HeartBeatReq{}, n.AppCenter, uint16(center.CMDCenter_IDHeartBeatReq), handleHeartBeatReq)
	g.EventRegister(g.ConnectSuccess, connectSuccess)
	g.EventRegister(g.Disconnect, disconnect)
	g.CallBackRegister(g.CbConfigChangeNotify, configChangeNotify)
}

func connectSuccess(args []interface{}) {
	a := args[g.AgentIndex].(n.AgentClient)
	log.Info("连接", "来了老弟,当前连接数=%d,name=%v", len(appConnData), a.AgentInfo().AppName)
	if v, ok := appConnData[a]; ok {
		log.Error("连接", "异常,重复连接?,%d,%d", v.appInfo.Type, v.appInfo.Id)
		a.Close()
		return
	}
	appConnData[a] = &connectionData{a: a, appInfo: lconf.BaseInfo{Name: a.AgentInfo().AppName}}
}

func disconnect(args []interface{}) {
	log.Info("连接", "告辞中,当前连接数=%d", len(appConnData))
	a := args[g.AgentIndex].(n.AgentClient)
	if v, ok := appConnData[a]; ok {
		regKey := util.MakeUint64FromUint32(v.appInfo.Type, v.appInfo.Id)
		log.Info("连接", "再见,appType=%d,appId=%d,regKey=%d", v.appInfo.Type, v.appInfo.Id, regKey)
		broadcastAppState(v.appInfo.Type, v.appInfo.Id, lconf.AppStateOffline)
		delete(appConnData, a)
		delete(appRegData, regKey)
	} else {
		log.Error("连接", "异常,没有注册的连接?")
	}
}

func configChangeNotify(args []interface{}) {
	key := args[apollo.KeyIndex].(apollo.ConfKey)
	//value := args[apollo.ValueIndex].(apollo.ConfValue)

	switch key.Key {
	case "服务列表":
		log.Debug("", "收到服务列表")
	default:
		break
	}
}

func handleRegisterAppReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*center.RegisterAppReq)
	a := args[n.AgentIndex].(n.AgentClient)

	//连接存在判断
	if _, ok := appConnData[a]; !ok {
		log.Error("连接", "异常,没有注册的连接?")
		a.Close()
		return
	}

	regKey := util.MakeUint64FromUint32(m.GetAppType(), m.GetAppId())
	if v, ok := appRegData[regKey]; ok {
		if v.regToken != m.GetReregToken() {

			resultMsg := fmt.Sprintf("该服务已注册,appType=%v,appId=%v,regKey=%v",
				m.GetAppType(), m.GetAppId(), regKey)
			log.Warning("连接", resultMsg)

			rsp := center.RegisterAppRsp{
				RegResult:  1,
				ReregToken: resultMsg,
				CenterId:   lconf.AppInfo.Id,
			}

			a.SendData(n.AppCenter, uint32(center.CMDCenter_IDAppRegRsp), &rsp)

			a.Close()
			return
		} else {
			log.Info("注册", "关闭旧的")
			v.a.Close()
		}
	} else {
		appRegData[regKey] = appConnData[a]
	}
	//信息存储
	token := fmt.Sprintf("gb%x%x%x", rand.Int(), time.Now().UnixNano(), rand.Int())
	appRegData[regKey].appInfo = lconf.BaseInfo{
		Type:         m.GetAppType(),
		Id:           m.GetAppId(),
		Name:         m.GetAppName(),
		ListenOnAddr: m.GetMyAddress(),
	}
	appRegData[regKey].regToken = token
	appRegData[regKey].appState = lconf.AppStateRunning

	log.Debug("注册", "服务注册,Name=%s,appType=%v,appId=%v,regKey=%v,addr=%v",
		m.GetAppName(), m.GetAppType(), m.GetAppId(), regKey, m.GetMyAddress())

	sendRsp := func(a n.AgentClient, i lconf.BaseInfo) {
		rsp := center.RegisterAppRsp{
			RegResult:  0,
			ReregToken: token,
			CenterId:   lconf.AppInfo.Id,
			AppName:    i.Name,
			AppType:    i.Type,
			AppId:      i.Id,
			AppAddress: i.ListenOnAddr,
		}
		log.Debug("注册", "发送注册回复,appInfo=%v", i)
		a.SendData(n.AppCenter, uint32(center.CMDCenter_IDAppRegRsp), &rsp)
	}

	//自己注册成功
	sendRsp(a, appRegData[regKey].appInfo)

	//daemon判断,只与配置中心连接
	if m.GetAppType() == n.AppDaemon {
		for k, v := range appRegData {
			if k == regKey || v.appInfo.Type != n.AppConfig {
				continue
			}
			//daemon判断
			sendRsp(a, v.appInfo)
			sendRsp(v.a, appRegData[regKey].appInfo)
		}
		return
	}

	//先广播已注册连接
	for k, v := range appRegData {
		if k == regKey || v.appInfo.Type == n.AppDaemon {
			continue
		}
		//sendRsp(a, v.appInfo)
		sendRsp(v.a, appRegData[regKey].appInfo)
	}

	//再通知当前
	for k, v := range appRegData {
		if k == regKey || v.appInfo.Type == n.AppDaemon {
			continue
		}
		sendRsp(a, v.appInfo)
		//sendRsp(v.a, appRegData[regKey].appInfo)
	}
}

func handleHeartBeatReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*center.HeartBeatReq)
	a := args[n.AgentIndex].(n.AgentClient)

	//非法判断
	if _, ok := appConnData[a]; !ok {
		log.Warning("心跳", "莫名的心跳?")
		return
	}
	app := appConnData[a]

	//log.Trace("", "心跳,aInfo=%v,state=%v,desc=%v,http=%v,rpc=%v",
	//	app.appInfo, m.GetServiceState(), m.GetStateDescription(), m.GetHttpAddress(), m.GetRpcAddress())

	app.totalHearDelayTime += time.Now().UnixNano() - m.GetPulseTime()
	app.lastHeartbeat = time.Now().UnixNano()
	if m.GetServiceState() != lconf.AppStateNone {
		app.appState = int(m.GetServiceState())
	}
	app.stateDescription = m.GetStateDescription()
	app.httpAddr = m.GetHttpAddress()
	app.rpcAddr = m.GetRpcAddress()

	//TODO 心跳回复,暂时无用
	//var rsp center.HeartBeatRsp
	//rsp.PulseTime = proto.Int64(time.Now().Unix())
	//a.SendData(n.AppCenter, uint32(center.CMDCenter_IDHeartBeatRsp), &rsp)
}

func broadcastAppState(appType, appId uint32, state int32) {
	for a, v := range appConnData {
		if v.appInfo.Type == appType && v.appInfo.Id == appId {
			continue
		}
		rsp := center.AppStateNotify{
			AppState: state,
			CenterId: lconf.AppInfo.Id,
			AppType:  appType,
			AppId:    appId,
		}
		a.SendData(n.AppCenter, uint32(center.CMDCenter_IDAppState), &rsp)
	}
}

type configServer struct {
	DaemonId uint32
	Alias    string
	lconf.BaseInfo
}

func getConfigServerList() []configServer {
	sList := make([]configServer, 0)
	v := apollo.GetConfig("服务列表", "")
	if err := json.Unmarshal([]byte(v), &sList); err != nil {
		log.Warning("", "反序列化服务列表出错,err=%v", err)
	}
	return sList
}

func getBaseInfoFromConfigList(appType, appId uint32) *configServer {
	sList := getConfigServerList()
	for _, v := range sList {
		if v.Type == appType && v.Id == appId {
			return &v
		}
	}
	return nil
}
