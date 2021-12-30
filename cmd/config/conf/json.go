package conf

import (
	"encoding/json"
	"io/ioutil"
	"mango/pkg/log"
	aConfig "mango/third_party/agollo/env/config"
)

var Server struct {
	CenterAddr    string
	ListenOnAddr  string
	AppID         uint32
	AppName       string
	FilePath      string
	ScreenPrint   bool
	UseApollo     bool `default:"false" json:"UseApollo"`
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
