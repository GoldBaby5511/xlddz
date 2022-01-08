package business

import (
	"bufio"
	"io"
	"mango/api/client"
	"mango/cmd/robot/business/player"
	"mango/pkg/conf/apollo"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/timer"
	"os"
	"strings"
	"time"
)

var (
	pls         []*player.Player = make([]*player.Player, 0)
	curWorkMode int64            = 0
)

func init() {
	g.MsgRegister(&client.LoginRsp{}, n.CMDClient, uint16(client.CMDID_Client_IDLoginRsp), handleLoginRsp)
	g.EventRegister(g.ConnectSuccess, connectSuccess)
	g.EventRegister(g.Disconnect, disconnect)
	g.EventRegister(g.ConfigChangeNotify, configChangeNotify)

	loadRobotAccounts()

	//g.Skeleton.LoopFunc(10*time.Second, checkRoomList, timer.LoopForever)
	//g.Skeleton.LoopFunc(1*time.Second, playerLogin, timer.LoopForever)
	//g.Skeleton.LoopFunc(1*time.Second, randJoinRoom, timer.LoopForever)
	//g.Skeleton.LoopFunc(1*time.Second, actionInRoom, timer.LoopForever)
}

func connectSuccess(args []interface{}) {
}

func disconnect(args []interface{}) {
}

func configChangeNotify(args []interface{}) {
	mode := apollo.GetConfigAsInt64("工作模式", 0)
	if mode != 0 && curWorkMode == 0 {
		curWorkMode = mode
		g.Skeleton.LoopFunc(1*time.Second, randJoinRoom, timer.LoopForever)
	}

}

func handleLoginRsp(args []interface{}) {
	//a := args[n.AgentIndex].(n.AgentClient)
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.LoginRsp)
	//srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	log.Debug("登录", "收到登录,主渠道=%d,UserId=%d", m.GetLoginResult(), m.GetBaseInfo().Account, m.GetBaseInfo().UserId)

	//sendLoginRsp(srcData.GetGateconnid(), "成功", uint32(client.LoginRsp_SUCCESS))
}

func checkRoomList() {
	loginCount := 0
	for _, pl := range pls {
		if pl.State < player.LoggedIn {
			continue
		}
		loginCount++
		if loginCount >= 50 {
			break
		}

		log.Debug("", "获取列表,a=%v,p=%v", pl.Account, pl.PassWord)
		pl.CheckRoomList()
	}
}

func playerLogin() {
	if len(pls) == 0 {
		return
	}
	loginCount := 0
	for _, pl := range pls {
		if pl.State != player.NilState {
			continue
		}
		loginCount++
		if loginCount >= 50 {
			break
		}
		pl.State = player.Logging

		log.Debug("", "登录,a=%v,p=%v", pl.Account, pl.PassWord)
		pl.Connect()
	}
}

func randJoinRoom() {

	loginCount := 0
	for _, pl := range pls {
		if pl.State != player.LoggedIn {
			continue
		}
		loginCount++
		if loginCount >= 50 {
			break
		}

		pl.JoinRoom()
	}

}

func actionInRoom() {

	loginCount := 0
	for _, pl := range pls {
		if pl.State != player.StandingInRoom {
			continue
		}
		loginCount++
		if loginCount >= 50 {
			break
		}

		pl.ActionRoom()
	}

}

func loadRobotAccounts() {
	f, err := os.Open("account.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')

		if err != nil || io.EOF == err {
			break
		}

		a := strings.Split(line[:len(line)-1], ",")
		pl := player.NewPlayer(a)
		if pl != nil {
			pls = append(pls, pl)
		}
	}

	log.Debug("", "pl=%d", len(pls))
}
