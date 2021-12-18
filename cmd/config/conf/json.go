package conf

import (
	"encoding/json"
	"io/ioutil"
	"xlddz/pkg/log"
	aConfig "xlddz/third_party/agollo/env/config"
)

var Server struct {
	CenterAddr      string
	ListenOnAddress string
	AppID           uint32
	AppName         string
	MaxConnNum      int
	FilePath        string
	ScreenPrint     bool
	UseApollo       bool `default:"false" json:"UseApollo"`
	Config          aConfig.AppConfig
	CommonServers   []ApolloConfig
}

type ApolloConfig struct {
	Appid      string `json:"appId"`
	Cluster    string `json:"cluster"`
	Ns         string `json:"namespaceName"`
	Ip         string `json:"ip"`
	ServerType uint32 `json:"servertype"`
	ServerId   uint32 `json:"serverid"`
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
