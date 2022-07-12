package business

import (
	"fmt"
	"mango/cmd/robot/business/player"
	"github.com/GoldBaby5511/go-mango-core/conf/apollo"
	g "github.com/GoldBaby5511/go-mango-core/gate"
	"github.com/GoldBaby5511/go-mango-core/log"
)

var (
	userList    []*player.Player = make([]*player.Player, 0)
	curWorkMode int64            = 0
)

func init() {
	g.EventRegister(g.ConfigChangeNotify, configChangeNotify)
}

func configChangeNotify(args []interface{}) {
	mode := apollo.GetConfigAsInt64("工作模式", 0)
	if mode != 0 && curWorkMode == 0 {
		curWorkMode = mode
		robotCount := apollo.GetConfigAsInt64("机器人数量", 10000)
		log.Debug("", "开始创建,robotCount=%v", robotCount)
		for i := 0; i < int(robotCount); i++ {
			pl := player.NewPlayer(fmt.Sprintf("robot%05d", i), "", 666)
			if pl != nil {
				userList = append(userList, pl)
			}
		}
	}
}
