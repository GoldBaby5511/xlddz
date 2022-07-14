package business

import (
	"fmt"
	"github.com/GoldBaby5511/go-mango-core/api/center"
	lconf "github.com/GoldBaby5511/go-mango-core/conf"
	g "github.com/GoldBaby5511/go-mango-core/gate"
	"github.com/GoldBaby5511/go-mango-core/log"
	n "github.com/GoldBaby5511/go-mango-core/network"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"time"
)

var (
	appConnData = make(map[n.AgentClient]*connectionData)
	appRegData  = make(map[uint64]*connectionData)
)

const (
	connected        int = 1
	registered       int = 2
	underMaintenance int = 3
)

//连接数据
type appRegInfo struct {
	appType   uint32
	appId     uint32
	regToken  string
	appName   string
	address   string
	curStatus int
}

type connectionData struct {
	a             n.AgentClient
	regInfo       appRegInfo
	lastHeartbeat int64
}

func init() {
	g.MsgRegister(&center.RegisterAppReq{}, n.AppCenter, uint16(center.CMDCenter_IDAppRegReq), handleRegisterAppReq)
	g.MsgRegister(&center.AppStateNotify{}, n.AppCenter, uint16(center.CMDCenter_IDAppState), handleAppStateNotify)
	g.MsgRegister(&center.HeartBeatReq{}, n.AppCenter, uint16(center.CMDCenter_IDHeartBeatReq), handleHeartBeatReq)
	g.EventRegister(g.ConnectSuccess, connectSuccess)
	g.EventRegister(g.Disconnect, disconnect)
	g.CallBackRegister(g.CbConfigChangeNotify, configChangeNotify)

	//apollo.RegPublicCB(configChangeNotify)
}

func connectSuccess(args []interface{}) {
	log.Info("连接", "来了老弟,当前连接数=%d", len(appConnData))
	a := args[g.AgentIndex].(n.AgentClient)
	if v, ok := appConnData[a]; ok {
		log.Error("连接", "异常,重复连接?,%d,%d", v.regInfo.appType, v.regInfo.appId)
		a.Close()
		return
	}
	appConnData[a] = &connectionData{a: a, regInfo: appRegInfo{curStatus: connected}}
}

func disconnect(args []interface{}) {
	log.Info("连接", "告辞中,当前连接数=%d", len(appConnData))
	a := args[g.AgentIndex].(n.AgentClient)
	if v, ok := appConnData[a]; ok {
		regKey := makeRegKey(v.regInfo.appType, v.regInfo.appId)
		log.Info("连接", "再见,appType=%d,appId=%d,regKey=%d", v.regInfo.appType, v.regInfo.appId, regKey)
		broadcastAppState(v.regInfo.appType, v.regInfo.appId)
		delete(appConnData, a)
		delete(appRegData, regKey)
	} else {
		log.Error("连接", "异常,没有注册的连接?")
	}
}

func configChangeNotify(args []interface{}) {
	//log.Info("配置", "真的收到了配置消息=%d", len(args))
	//key := apollo.ConfKey{AppType: lconf.AppInfo.Type, AppId: lconf.AppInfo.Id, Key: "服务维护"}
	//if k == key {
	//	type appInfo struct {
	//		AppType uint32
	//		AppId   uint32
	//		OpType  uint32
	//	}
	//	var info appInfo
	//	err := json.Unmarshal([]byte(v.Value), &info)
	//	if err != nil {
	//		log.Error("配置", "%v", err)
	//		return
	//	}
	//
	//	if _, ok := appRegData[makeRegKey(info.AppType, info.AppId)]; !ok {
	//		log.Warning("配置", "要维护的服务不存在啊,info=%v", info)
	//		return
	//	}
	//	appRegData[makeRegKey(info.AppType, info.AppId)].regInfo.curStatus = int(info.OpType)
	//
	//	log.Debug("配置", "收到服务维护配置,%v", info)
	//}
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

	//是否已注册
	if appConnData[a].regInfo.curStatus == registered {
		return
	}

	regKey := makeRegKey(m.GetAppType(), m.GetAppId())
	if v, ok := appRegData[regKey]; ok {
		if v.regInfo.regToken != m.GetReregToken() {

			resultMsg := fmt.Sprintf("该服务已注册,appType=%v,appId=%v,regKey=%v",
				m.GetAppType(), m.GetAppId(), regKey)
			log.Warning("连接", resultMsg)

			var rsp center.RegisterAppRsp
			rsp.RegResult = proto.Uint32(1)
			rsp.ReregToken = proto.String(resultMsg)
			rsp.CenterId = proto.Uint32(lconf.AppInfo.Id)
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
	appRegData[regKey].regInfo = appRegInfo{m.GetAppType(), m.GetAppId(), token, m.GetAppName(), m.GetMyAddress(), registered}

	log.Debug("注册", "服务注册,appType=%v,appId=%v,regKey=%v,%v",
		m.GetAppType(), m.GetAppId(), regKey, m.GetMyAddress())

	sendRsp := func(a n.AgentClient, i appRegInfo) {
		var rsp center.RegisterAppRsp
		rsp.RegResult = proto.Uint32(0)
		rsp.ReregToken = proto.String(token)
		rsp.CenterId = proto.Uint32(lconf.AppInfo.Id)
		rsp.AppName = proto.String(i.appName)
		rsp.AppType = proto.Uint32(i.appType)
		rsp.AppId = proto.Uint32(i.appId)
		rsp.AppAddress = proto.String(i.address)
		a.SendData(n.AppCenter, uint32(center.CMDCenter_IDAppRegRsp), &rsp)
	}

	//广播已注册
	for k, v := range appConnData {
		sendRsp(k, appRegData[regKey].regInfo)
		if v.regInfo.appType == m.GetAppType() && v.regInfo.appId == m.GetAppId() {
			continue
		}
		sendRsp(a, v.regInfo)
	}
}

func handleAppStateNotify(args []interface{}) {

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

	log.Trace("", "心跳,aInfo=%v,state=%v,desc=%v,http=%v,rpc=%v",
		app.lastHeartbeat, m.GetServiceState(), m.GetStateDescription(), m.GetHttpAddress(), m.GetRpcAddress())

	app.lastHeartbeat = time.Now().UnixNano()

	var rsp center.HeartBeatRsp
	rsp.PulseTime = proto.Int64(time.Now().Unix())
	a.SendData(n.AppCenter, uint32(center.CMDCenter_IDHeartBeatRsp), &rsp)

}

func makeRegKey(appType, appId uint32) uint64 {
	return uint64(appType)<<32 | uint64(appId)
}

func broadcastAppState(appType, appId uint32) {
	for a, v := range appConnData {
		if v.regInfo.appType == appType && v.regInfo.appId == appId {
			continue
		}
		var rsp center.AppStateNotify
		rsp.CenterId = proto.Uint32(lconf.AppInfo.Id)
		rsp.AppType = proto.Uint32(appType)
		rsp.AppId = proto.Uint32(appId)
		a.SendData(n.AppCenter, uint32(center.CMDCenter_IDAppState), &rsp)
	}
}
