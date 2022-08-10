package core

import (
	"errors"
	"featherOne/api"
	"featherOne/conf"
	"featherOne/core/utils"
	"github.com/projectdiscovery/gologger"
	"os"
)

type Runner struct {
	options     *utils.Options
	conf        *conf.Conf
	readyFofa   api.Apier
	readyQuake  api.Apier
	readyHunter api.Apier
}

func NewRunner(options *utils.Options) *Runner {
	return &Runner{
		options: options,
		conf:    conf.InitConf(),
	}
}

func (r *Runner) check() error {
	var apiObj api.Apier
	if r.options.SearchString != "" {
		if r.options.QuakeSearch {
			apiObj = api.NewQuake(r.conf.QuakeToken, r.options.Num)
			if ok := apiObj.Auth(); !ok {
				return api.NewApiError("Quake")
			}
			r.readyHunter = apiObj
		}
		if r.options.HunterSearch {
			apiObj = api.NewHunter(r.conf.HunterApiKey, r.options.Num)
			if ok := apiObj.Auth(); !ok {
				return api.NewApiError("Hunuter")
			}
			r.readyHunter = apiObj
		}
		return errors.New("please choose a search engine")
	}
	return nil
}

func (r *Runner) Search() {
	var resultDomain api.IPLists
	var tmpDomain api.IPLists

	err := r.check()
	if err != nil {
		gologger.Error().Msgf("Api auth err :%s", err)
		os.Exit(0)
	}

	if r.readyQuake != nil {
		result := api.Search(r.readyQuake, r.options.SearchString)
		tmpDomain = append(tmpDomain, result...)
	}
	if r.readyHunter != nil {
		result := api.Search(r.readyQuake, r.options.SearchString)
		tmpDomain = append(tmpDomain, result...)
	}
	if r.readyHunter != nil {
		result := api.Search(r.readyQuake, r.options.SearchString)
		tmpDomain = append(tmpDomain, result...)
	}

	if r.options.SameIp {
		result := api.SearchSip(r.options.Url)
		tmpDomain = append(tmpDomain, result...)
	}
	if r.options.Ip {
		result := api.SearchIp(r.options.Url)
		gologger.Silent().Msg(result)
		return
	}
	if r.options.Weight {
		w, err := api.SearchWeight(r.options.Url)
		if err != nil {
			gologger.Error().Msgf("%s", err)
			return
		}
		gologger.Silent().Msgf("%s baidu' weight is %v", r.options.Url, w)
		return
	}

	// Unique results
	tmpList := make(map[string]int)
	for _, item := range tmpDomain {
		if item == "" {
			continue
		}
		if _, ok := tmpList[item]; !ok {
			gologger.Silent().Msg(item)
			resultDomain = append(resultDomain, item)
			tmpList[item] = 1
		}
	}

	gologger.Info().Msgf("Total target url number is %v\n", len(resultDomain))

	//if !r.options.Silent {
	//	pool := utils.NewPool(resultDomain)
	//	pool.Start()
	//}

}
