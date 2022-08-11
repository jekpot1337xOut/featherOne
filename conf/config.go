package conf

import (
	"featherOne/core/utils"

	"github.com/projectdiscovery/gologger"
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
	confP := utils.GetConfPath()
	gologger.Info().Msgf("Load config file %s\n", confP)
	viper.SetConfigFile(confP)
	err := viper.ReadInConfig()
	if err != nil {
		gologger.Error().Msgf("Fatal error config file: %s \n", err)
	}
	conf.FofaEmail = viper.GetStringMapString("fofa")["email"]
	conf.FofaToken = viper.GetStringMapString("fofa")["fofatoken"]
	conf.QuakeToken = viper.GetStringMapString("quake")["quaketoken"]
	conf.DnsgrepToken = viper.GetStringMapString("dnsgrep")["dnsgreptoken"]
	conf.HunterApiKey = viper.GetStringMapString("hunter")["hunterapikey"]

	return conf
}
