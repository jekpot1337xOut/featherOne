package conf

import (
	"fmt"

	"featherOne/core/utils"

	"github.com/spf13/viper"
)

type Conf struct {
	FofaEmail    string
	FofaToken    string
	QuakeToken   string
	DnsgrepToken string
	HunterApiKey string
}

func InitConf() *Conf {
	conf := &Conf{}
	viper.AddConfigPath("F:\\language\\golang_folder\\featherOne\\conf\\")
	//viper.SetConfigFile("F:\\language\\golang_folder\\featherOne\\conf\\config.yaml")
	viper.SetConfigFile(utils.GetConfPath())
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	err := viper.ReadInConfig()   // 查找并读取配置文件
	if err != nil {               // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	conf.FofaEmail = viper.GetStringMapString("fofa")["email"]
	conf.FofaToken = viper.GetStringMapString("fofa")["fofatoken"]
	conf.QuakeToken = viper.GetStringMapString("quake")["quaketoken"]
	conf.DnsgrepToken = viper.GetStringMapString("dnsgrep")["dnsgreptoken"]
	conf.HunterApiKey = viper.GetStringMapString("hunter")["hunterapikey"]

	return conf
}
