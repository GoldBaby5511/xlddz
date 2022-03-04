package gate

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"mango/api/center"
	"mango/api/config"
	"mango/api/gateway"
	"mango/api/logger"
	"mango/pkg/chanrpc"
	"mango/pkg/conf"
	"mango/pkg/conf/apollo"
	"mango/pkg/database"
	"mango/pkg/log"
	"mango/pkg/module"
	n "mango/pkg/network"
	"mango/pkg/network/protobuf"
	"mango/pkg/util"

	"os"
	"os/signal"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	//事件
	ConnectSuccess     string = "ConnectSuccess"
	Disconnect         string = "Disconnect"
	ConfigChangeNotify string = "ConfigChangeNotify"
	CenterConnected    string = "CenterConnected"
	CenterDisconnect   string = "CenterDisconnect"
	CenterRegResult    string = "CenterRegResult"
	CommonServerReg    string = "CommonServerReg"

	AgentIndex = 0
)

var (
	cbCenterDisconnect []func()
	tcpLog             *n.TCPClient = nil
	mxServers          sync.Mutex
	wg                 sync.WaitGroup
	servers            map[uint64]*agentServer = make(map[uint64]*agentServer)
	agentChanRPC       *chanrpc.Server         = nil
	Skeleton           *module.Skeleton        = nil
	processor                                  = protobuf.NewProcessor()
	MaxConnNum         int
	PendingWriteNum    int
	MinMsgLen          uint32
	MaxMsgLen          uint32

	// websocket
	WSAddr      string
	HTTPTimeout time.Duration
	CertFile    string
	KeyFile     string

	closeSig chan bool
)

func init() {
	tcpLog = new(n.TCPClient)
	cbCenterDisconnect = append(cbCenterDisconnect, apollo.CenterDisconnect)
	Skeleton = module.NewSkeleton(conf.GoLen, conf.TimerDispatcherLen, conf.AsynCallLen, conf.ChanRPCLen)
	agentChanRPC = Skeleton.ChanRPCServer
	closeSig = make(chan bool, 0)
	MsgRegister(&config.ApolloCfgRsp{}, n.AppConfig, uint16(config.CMDConfig_IDApolloCfgRsp), handleApolloCfgRsp)
}

func Start(appName string) {
	conf.AppInfo.Name = appName
	// logger
	l, err := log.New(conf.AppInfo.Name)
	if err != nil {
		panic(err)
	}
	log.Export(l)
	defer l.Close()

	//baseConfig
	conf.LoadBaseConfig()

	if conf.AppInfo.Type == n.AppCenter {
		apollo.RegisterConfig("", conf.AppInfo.Type, conf.AppInfo.Id, nil)
	}

	wg.Add(2)
	go func() {
		Skeleton.Run()
		wg.Done()
	}()

	go func() {
		Run()
		wg.Done()
	}()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Info("主流程", "服务器关闭 (signal: %v)", sig)
	Stop()
}

func Stop() {
	defer util.TryE(conf.AppInfo.Name)
	for k, v := range servers {
		v.Close()
		delete(servers, k)
	}
	closeSig <- true
	time.Sleep(time.Second / 2)
	Skeleton.Close()
	wg.Wait()
}

func MsgRegister(m proto.Message, appType uint32, cmdId uint16, f interface{}) {
	chanRPC := Skeleton.ChanRPCServer
	processor.Register(m, appType, cmdId, chanRPC)
	chanRPC.Register(reflect.TypeOf(m), f)
}

func EventRegister(id interface{}, f interface{}) {
	Skeleton.ChanRPCServer.Register(id, f)
}

func Run() {
	log.Debug("", "Run,ListenOnAddr=%v", conf.AppInfo.ListenOnAddr)

	var wsServer *n.WSServer
	if WSAddr != "" {
		wsServer = new(n.WSServer)
		wsServer.Addr = WSAddr
		wsServer.MaxConnNum = MaxConnNum
		wsServer.PendingWriteNum = PendingWriteNum
		wsServer.MaxMsgLen = MaxMsgLen
		wsServer.HTTPTimeout = HTTPTimeout
		wsServer.CertFile = CertFile
		wsServer.KeyFile = KeyFile
		wsServer.NewAgent = func(conn *n.WSConn) n.AgentClient {
			a := &agentClient{conn: conn}
			if agentChanRPC != nil {
				agentChanRPC.Go(ConnectSuccess, a)
			}
			return a
		}
	}

	var tcpServer *n.TCPServer
	if conf.AppInfo.ListenOnAddr != "" {
		tcpServer = new(n.TCPServer)
		tcpServer.Addr = conf.AppInfo.ListenOnAddr
		tcpServer.MaxConnNum = MaxConnNum
		tcpServer.PendingWriteNum = PendingWriteNum
		tcpServer.MinMsgLen = MinMsgLen
		tcpServer.MaxMsgLen = MaxMsgLen
		tcpServer.GetConfig = apollo.GetConfigAsInt64
		tcpServer.NewAgent = func(conn *n.TCPConn) n.AgentClient {
			a := &agentClient{conn: conn, info: n.BaseAgentInfo{AgentType: n.NormalUser}}
			if agentChanRPC != nil {
				agentChanRPC.Go(ConnectSuccess, a)
			}
			return a
		}
	}

	if conf.AppInfo.CenterAddr != "" && conf.AppInfo.Type != n.AppCenter && conf.AppInfo.Type != n.AppLogger {
		newServerItem(n.BaseAgentInfo{AgentType: n.CommonServer, AppName: "center", AppType: n.AppCenter, ListenOnAddr: conf.AppInfo.CenterAddr}, true, PendingWriteNum)
	}

	if wsServer != nil {
		wsServer.Start()
	}
	if tcpServer != nil {
		tcpServer.Start()
	}

	<-closeSig
	if wsServer != nil {
		wsServer.Close()
	}
	if tcpServer != nil {
		tcpServer.Close()
	}
	if tcpLog != nil {
		tcpLog.Close()
	}
}

func handleApolloCfgRsp(args []interface{}) {
	apollo.ProcessConfigRsp(args[n.DataIndex].(n.BaseMessage).MyMessage.(*config.ApolloCfgRsp))

	dbConfig := apollo.GetConfig("数据库配置", "")
	if database.DBC == nil && dbConfig != "" {
		database.InitDBHelper(dbConfig)
	}

	logAddr := apollo.GetConfig("日志服务器地址", "")
	if logAddr != "" && tcpLog != nil && !tcpLog.IsRunning() {
		ConnectLogServer(logAddr)
	}

	go func() {
		if agentChanRPC != nil {
			agentChanRPC.Call0(ConfigChangeNotify)
		}
	}()
}

func ConnectLogServer(logAddr string) {
	log.Info("gate", "连接日志服务器,Addr=%v", logAddr)
	if conf.AppInfo.Type != n.AppLogger && logAddr != "" && tcpLog != nil && !tcpLog.IsRunning() {
		if conf.RunInLocalDocker() {
			addr := strings.Split(logAddr, "|")
			logAddr = ""
			for i, v := range addr {
				curConnAddr := v
				logAddr = logAddr + "logger:" + strconv.Itoa(util.GetPortFromIPAddress(curConnAddr))
				if i != len(addr)-1 {
					logAddr = logAddr + "|"
				}
			}
		}
		tcpLog.Addr = logAddr
		tcpLog.AutoReconnect = true
		tcpLog.NewAgent = func(conn *n.TCPConn) n.AgentServer {
			a := &agentServer{tcpClient: tcpLog, conn: conn, info: n.BaseAgentInfo{AgentType: n.CommonServer, AppName: "logger", AppType: n.AppLogger, AppID: 0, ListenOnAddr: logAddr}}
			log.Info("gate", "日志服务器连接成功")

			log.SetCallback(func(i log.LogInfo) {
				var logReq logger.LogReq
				logReq.FileName = proto.String(i.File)
				logReq.LineNo = proto.Uint32(uint32(i.Line))
				logReq.SrcApptype = proto.Uint32(conf.AppInfo.Type)
				logReq.SrcAppid = proto.Uint32(conf.AppInfo.Id)
				logReq.Content = []byte(i.LogStr)
				logReq.ClassName = []byte(i.Classname)
				logReq.LogLevel = proto.Uint32(uint32(i.Level))
				logReq.TimeNs = proto.Int64(i.TimeNs)
				logReq.SrcAppname = proto.String(conf.AppInfo.Name)
				cmd := n.TCPCommand{AppType: uint16(n.AppLogger), CmdId: uint16(logger.CMDLogger_IDLogReq)}
				bm := n.BaseMessage{MyMessage: &logReq, Cmd: cmd}
				a.SendMessage(bm)
			})

			return a
		}

		tcpLog.Start()
	}
}

func sendRegAppReq(a *agentServer) {
	var registerReq center.RegisterAppReq
	registerReq.AuthKey = proto.String("GoldBaby")
	registerReq.AppName = proto.String(conf.AppInfo.Name)
	registerReq.AppType = proto.Uint32(conf.AppInfo.Type)
	registerReq.AppId = proto.Uint32(conf.AppInfo.Id)
	myAddress := conf.AppInfo.ListenOnAddr
	if conf.RunInLocalDocker() {
		myAddress = conf.AppInfo.Name + ":" + strconv.Itoa(util.GetPortFromIPAddress(conf.AppInfo.ListenOnAddr))
	}
	registerReq.MyAddress = proto.String(myAddress)
	a.SendData(n.AppCenter, uint32(center.CMDCenter_IDAppRegReq), &registerReq)
}

func SendData(dataSrc n.BaseAgentInfo, bm n.BaseMessage) error {
	if dataSrc.AgentType == n.CommonServer {
		return sendData(bm, dataSrc.AppType, dataSrc.AppID)
	}
	return SendMessage2Client(bm, util.MakeUint64FromUint32(dataSrc.AppType, dataSrc.AppID), 0)
}

func SendData2App(destAppType, destAppid, appType, cmdId uint32, m proto.Message) error {
	cmd := n.TCPCommand{AppType: uint16(appType), CmdId: uint16(cmdId)}
	bm := n.BaseMessage{MyMessage: m, Cmd: cmd}
	return sendData(bm, destAppType, destAppid)
}

func SendMessage2Client(bm n.BaseMessage, gateConnID, sessionID uint64) error {
	var dataReq gateway.TransferDataReq
	dataReq.DestApptype = proto.Uint32(n.AppGate)
	dataReq.DestAppid = proto.Uint32(util.GetLUint32FromUint64(gateConnID))
	dataReq.DataApptype = proto.Uint32(uint32(bm.Cmd.AppType))
	dataReq.DataCmdid = proto.Uint32(uint32(bm.Cmd.CmdId))
	dataReq.Data, _ = proto.Marshal(bm.MyMessage.(proto.Message))
	dataReq.Gateconnid = proto.Uint64(gateConnID)
	dataReq.AttSessionid = proto.Uint64(sessionID)
	cmd := n.TCPCommand{AppType: uint16(n.AppGate), CmdId: uint16(gateway.CMDGateway_IDTransferDataReq)}
	transBM := n.BaseMessage{MyMessage: &dataReq, Cmd: cmd, TraceId: bm.TraceId}
	return sendData(transBM, n.AppGate, util.GetLUint32FromUint64(gateConnID))
}

func sendData(bm n.BaseMessage, destAppType, destAppid uint32) error {
	destAgents := getDestAppInfo(destAppType, destAppid)
	if len(destAgents) == 0 {
		return fmt.Errorf("目标没找到,destAppType=%d,destAppid=%d", destAppType, destAppid)
	}
	for _, a := range destAgents {
		a.SendMessage(bm)
	}
	return nil
}

func getDestAppInfo(destAppType, destAppid uint32) []*agentServer {
	mxServers.Lock()
	defer mxServers.Unlock()
	var destAgent []*agentServer
	destTypeAppCount := func() int {
		destCount := 0
		for _, v := range servers {
			if v.info.AppType == destAppType {
				destCount++
			}
		}
		return destCount
	}
	if destTypeAppCount() != 0 {
		switch destAppid {
		case n.Send2All:
			for _, v := range servers {
				if v.info.AppType == destAppType {
					destAgent = append(destAgent, v)
				}
			}
		case n.Send2AnyOne:
			for _, v := range servers {
				if v.info.AppType == destAppType {
					destAgent = append(destAgent, v)
					break
				}
			}
		default:
			for _, v := range servers {
				if v.info.AppType == destAppType && v.info.AppID == destAppid {
					destAgent = append(destAgent, v)
					break
				}
			}
		}
	}

	if len(destAgent) == 0 {
		log.Error("转发", "异常,消息转发失败,appCount=%v,destAppType=%v,destAppid=%v",
			destTypeAppCount(), destAppType, destAppid)
	}

	return destAgent
}
