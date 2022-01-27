package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
	"strconv"
)

const (
	ArgAppName      string = "-Name"
	ArgAppType      string = "-Type"
	ArgAppId        string = "-Id"
	ArgCenterAddr   string = "-CenterAddr"
	ArgListenOnAddr string = "-ListenOnAddr"
	ArgDockerRun    string = "-DockerRun"
)

var (
	LenStackBuf        = 4096
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
	AppInfo            BaseInfo
)

type BaseInfo struct {
	Name         string
	Type         uint32
	Id           uint32
	ListenOnAddr string
	CenterAddr   string
}

func LoadBaseConfig() {
	if AppInfo.Name != "" {
		data, err := ioutil.ReadFile(fmt.Sprintf("configs/%s/%s.json", AppInfo.Name, AppInfo.Name))
		if err == nil {
			err = json.Unmarshal(data, &AppInfo)
		}
	}

	if v, ok := util.ParseArgsString(ArgAppName); ok {
		AppInfo.Name = v
	}
	if v, ok := util.ParseArgsUint32(ArgAppType); ok {
		AppInfo.Type = v
	}
	if v, ok := util.ParseArgsUint32(ArgAppId); ok {
		AppInfo.Id = v
	}
	if v, ok := util.ParseArgsString(ArgCenterAddr); ok {
		AppInfo.CenterAddr = v
	}
	if v, ok := util.ParseArgsString(ArgListenOnAddr); ok {
		AppInfo.ListenOnAddr = v
	}
	if AppInfo.ListenOnAddr == "" {
		AppInfo.ListenOnAddr = fmt.Sprintf("0.0.0.0:%d", 10000+AppInfo.Id)
	}
	if AppInfo.CenterAddr == "" && AppInfo.Type != n.AppCenter {
		AppInfo.CenterAddr = "127.0.0.1:10050"
		log.Debug("", "未指定中心服,使用默认地址,CenterAddr=%v", AppInfo.CenterAddr)
	}
	if RunInLocalDocker() {
		AppInfo.CenterAddr = "center:" + strconv.Itoa(util.GetPortFromIPAddress(AppInfo.CenterAddr))
	}

	portPID := util.PortInUse(util.GetPortFromIPAddress(AppInfo.ListenOnAddr))
	if portPID != -1 {
		log.Fatal("初始化", "端口[%v]已被[PID=%v]占用,请检查运行环境", util.GetPortFromIPAddress(AppInfo.ListenOnAddr), portPID)
		return
	}

	if AppInfo.Name == "" || AppInfo.Type == 0 || AppInfo.Id == 0 || AppInfo.ListenOnAddr == "" ||
		(AppInfo.CenterAddr == "" && AppInfo.Type != n.AppCenter) {
		log.Fatal("初始化", "初始参数异常,请检查,AppInfo=%v", AppInfo)
		return
	}

	log.Debug("", "基础属性,%v", AppInfo)
}

func RunInLocalDocker() bool {
	if v, ok := util.ParseArgsUint32(ArgDockerRun); ok && v == 1 {
		return true
	}
	return false
}
