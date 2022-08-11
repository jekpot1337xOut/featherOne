package core

import (
	"errors"
	"os"

	"featherOne/api"
	"featherOne/conf"
	"featherOne/core/utils"

	"github.com/projectdiscovery/gologger"
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
			gologger.Info().Msgf("Check quake authorization,wait a second... ")
			apiObj = api.NewQuake(r.conf.QuakeToken, r.options.Num)
			if ok := apiObj.Auth(); !ok {
				return api.NewApiError("Quake")
			}
			r.readyQuake = apiObj
		} else if r.options.HunterSearch {
			gologger.Info().Msgf("Check hunter authorization,wait a second...\n")
			apiObj = api.NewHunter(r.conf.HunterApiKey, r.options.Num)
			if ok := apiObj.Auth(); !ok {
				return api.NewApiError("Hunuter")
			}
			r.readyHunter = apiObj
		} else {
			return errors.New("please choose a search engine")
		}
	}
	return nil
}

func (r *Runner) Search() {
	var resultDomain api.IPLists
	var tmpDomain api.IPLists

	err := r.check()
	if err != nil {
		gologger.Error().Msgf("Api auth err :%s\n", err)
		os.Exit(0)
	}

	gologger.Info().Msgf("Search grammar is: %s\n", r.options.SearchString)

	// TODO
	// add new sentence if you add new api
	if r.readyQuake != nil {
		gologger.Info().Msg("Searching data via quake...")
		if r.options.AutoGrammar {
			r.options.SearchString = api.AutoGrammer(r.options.SearchString, "quake")
			gologger.Info().Msgf("transfer grammar is %s\n", r.options.SearchString)
		}
		result := api.Search(r.readyQuake, r.options.SearchString)
		tmpDomain = append(tmpDomain, result...)
	}
	if r.readyHunter != nil {
		gologger.Info().Msg("Searching data via hunter...")
		if r.options.AutoGrammar {
			r.options.SearchString = api.AutoGrammer(r.options.SearchString, "hunter")
			gologger.Info().Msgf("transfer grammar is %s\n", r.options.SearchString)
		}
		result := api.Search(r.readyHunter, r.options.SearchString)
		tmpDomain = append(tmpDomain, result...)
	}

	if r.options.SameIp {
		gologger.Info().Msg("Searching same ip...")
		result := api.SearchSip(r.options.Url)
		tmpDomain = append(tmpDomain, result...)
	}
	if r.options.Ip {
		gologger.Info().Msg("Searching ip...")
		result := api.SearchIp(r.options.Url)
		gologger.Silent().Msg(result)
		return
	}
	if r.options.Weight {
		gologger.Info().Msg("Searching weight...")
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

	if r.options.Probe {
		pool := utils.NewPool(resultDomain)
		pool.Start()
	}

}
