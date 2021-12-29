package apollo

import (
	"github.com/golang/protobuf/proto"
	"strconv"
	"sync"
	"time"
	"mango/api/config"
	"mango/pkg/conf"
	"mango/pkg/log"
	"mango/pkg/network"
)

var (
	configValues map[ConfKey]*ConfValue = make(map[ConfKey]*ConfValue)
	regSubList   map[ConfKey]*ConfValue = make(map[ConfKey]*ConfValue)
	netAgent     network.AgentServer    = nil
	mutexConfig  sync.Mutex
	mxRegSub     sync.Mutex
)

type cbNotify func(key ConfKey, value ConfValue)

type ConfKey struct {
	AppType uint32
	AppId   uint32
	Key     string
}

type ConfValue struct {
	Value    string
	RspCount uint64
	Cb       cbNotify
}

func init() {
	log.GetMinLevelConfig = GetConfigAsInt64
}

func SetNetAgent(a network.AgentServer) {
	netAgent = a

	//连接成功后开启定时订阅
	timeInterval := 30 * time.Second
	timer := time.NewTimer(timeInterval)
	go func(t *time.Timer) {
		for {
			<-t.C

			//保持订阅
			for key, _ := range regSubList {
				SendSubscribeReq(key, false)
			}

			t.Reset(timeInterval)
		}
	}(timer)
}

func CenterDisconnect() {
	regSubList = make(map[ConfKey]*ConfValue)
}

func ProcessConfigRsp(m *config.ApolloCfgRsp) {
	if len(m.GetKey()) == 0 || (len(m.GetKey()) != len(m.GetValue())) {
		log.Error("apollo", "异常,收到空的Apollo配置,PacketId=%v,ns=%v,key=%v,type=%v,appid=%v",
			m.GetPacketId(), m.GetNameSpace(), m.GetKeyName(), m.GetSubAppType(), m.GetSubAppId())
		return
	}

	nsKey := ConfKey{Key: m.GetKeyName(), AppType: m.GetSubAppType(), AppId: m.GetSubAppId()}
	mxRegSub.Lock()
	if _, ok := regSubList[nsKey]; !ok {
		mxRegSub.Unlock()
		log.Error("apollo", "异常，返回的竟然是自己没订阅的")
		return
	}
	regSubList[nsKey].RspCount += 1
	mxRegSub.Unlock()

	for i := 0; i < len(m.GetKey()); i++ {
		key := ConfKey{Key: m.GetKey()[i], AppType: m.GetSubAppType(), AppId: m.GetSubAppId()}
		if m.GetKey()[i] == "LogScreenPrint" {
			p, _ := strconv.Atoi(m.GetValue()[i])
			log.SetScreenPrint(p)
		}
		mutexConfig.Lock()
		if _, ok := configValues[key]; ok {
			configValues[key].Value = m.GetValue()[i]
			configValues[key].RspCount += 1
		} else {
			configValues[key] = &ConfValue{Value: m.GetValue()[i], RspCount: 1}
		}
		mutexConfig.Unlock()
	}
}

func GetConfig(key, defaultValue string) string {
	nsKey := ConfKey{Key: key, AppType: conf.AppInfo.AppType, AppId: conf.AppInfo.AppID}
	mutexConfig.Lock()
	defer mutexConfig.Unlock()
	if item, ok := configValues[nsKey]; ok {
		return item.Value
	}
	return defaultValue
}

func GetConfigAsInt64(key string, defaultValue int64) int64 {
	v, _ := strconv.ParseInt(GetConfig(key, strconv.FormatInt(defaultValue, 10)), 10, 64)
	return v
}

func RegisterConfig(key string, reqAppType, reqAppId uint32, cb cbNotify) {
	mxRegSub.Lock()
	nsKey := ConfKey{Key: key, AppType: reqAppType, AppId: reqAppId}
	if _, ok := regSubList[nsKey]; ok {
		log.Info("Apollo", "这个key已经注册过了")
		return
	}

	regSubList[nsKey] = &ConfValue{Cb: cb}
	mxRegSub.Unlock()
	log.Info("Apollo", "注册Apollo订阅，%v", nsKey)

	SendSubscribeReq(nsKey, false)
}

func SendSubscribeReq(k ConfKey, cancel bool) {
	if netAgent == nil {
		return
	}
	mxRegSub.Lock()
	defer mxRegSub.Unlock()
	if _, ok := regSubList[k]; !ok {
		return
	}

	var req config.ApolloCfgReq
	req.AppType = proto.Uint32(conf.AppInfo.AppType)
	req.AppId = proto.Uint32(conf.AppInfo.AppID)
	req.SubAppType = proto.Uint32(k.AppType)
	req.SubAppId = proto.Uint32(k.AppId)
	req.KeyName = proto.String(k.Key)
	subscribe := config.ApolloCfgReq_SUBSCRIBE
	if regSubList[k].RspCount == 0 {
		subscribe = subscribe | config.ApolloCfgReq_NEED_RSP
	}
	if cancel {
		subscribe = config.ApolloCfgReq_UNSUBSCRIBE
	}
	req.Subscribe = proto.Uint32(uint32(subscribe))

	cmd := network.TCPCommand{MainCmdID: uint16(network.AppConfig), SubCmdID: uint16(config.CMDID_Config_IDApolloCfgReq)}
	bm := network.BaseMessage{MyMessage: &req, Cmd: cmd}
	netAgent.SendMessage(bm)
}
