package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mango/pkg/log"
	"mango/pkg/util"
	"strings"
)

const (
	ArgAppName      string = "/AppName"
	ArgAppType      string = "/AppType"
	ArgAppID        string = "/AppID"
	ArgCenterAddr   string = "/CenterAddr"
	ArgListenOnAddr string = "/ListenOnAddr"
	ArgDockerRun    string = "/DockerRun"
)

var (
	LenStackBuf = 4096

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
	AppInfo            BaseInfo
)

type BaseInfo struct {
	AppName      string
	AppType      uint32
	AppID        uint32
	ListenOnAddr string
	CenterAddr   string
}

func LoadBaseConfig() {
	if AppInfo.AppName != "" {
		data, err := ioutil.ReadFile(fmt.Sprintf("configs/%s/%s.json", AppInfo.AppName, AppInfo.AppName))
		if err == nil {
			err = json.Unmarshal(data, &AppInfo)
		}
	}

	if v, ok := util.ParseArgsString(ArgAppName); ok {
		AppInfo.AppName = v
	}
	if v, ok := util.ParseArgsUint32(ArgAppType); ok {
		AppInfo.AppType = v
	}
	if v, ok := util.ParseArgsUint32(ArgAppID); ok {
		AppInfo.AppID = v
	}
	if v, ok := util.ParseArgsString(ArgCenterAddr); ok {
		AppInfo.CenterAddr = v
	}
	if v, ok := util.ParseArgsString(ArgListenOnAddr); ok {
		AppInfo.ListenOnAddr = v
	}
	if v, ok := util.ParseArgsUint32(ArgDockerRun); ok && v == 1 {
		addr := strings.Split(AppInfo.CenterAddr, ":")
		if len(addr) == 2 {
			AppInfo.CenterAddr = "center:" + addr[1]
		}
	}

	if AppInfo.AppName == "" || AppInfo.AppType == 0 || AppInfo.AppID == 0 {
		log.Fatal("初始化", "初始参数异常,请检查,AppInfo=%v", AppInfo)
	}

	log.Debug("", "参数解析,%v", AppInfo)
}
