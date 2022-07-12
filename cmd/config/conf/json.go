package conf

import (
	"encoding/json"
	"io/ioutil"
	"github.com/GoldBaby5511/go-mango-core/log"
	aConfig "mango/third_party/agollo/env/config"
)

var Server struct {
	AppName      string
	AppType      uint32
	AppID        uint32
	ListenOnAddr string
	CenterAddr   string

	UseApollo     bool `default:"false" json:"UseApollo"`
	LoggerAddr    string
	Config        aConfig.AppConfig
	CommonServers []ApolloConfig
}

type ApolloConfig struct {
	Appid      string `json:"appID"`
	Cluster    string `json:"cluster"`
	Ns         string `json:"namespaceName"`
	Ip         string `json:"ip"`
	ServerType uint32 `json:"serverType"`
	ServerId   uint32 `json:"serverID"`
}

func init() {
	data, err := ioutil.ReadFile("configs/config/config.json")
	if err != nil {
		log.Fatal("jsonconf", "%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("jsonconf", "%v", err)
	}

	log.Info("jsonconf", "配置文件载入成功%v", Server)
}
