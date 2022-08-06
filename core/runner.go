package core

import (
	"featherOne/api"
	"featherOne/conf"
	"featherOne/core/utils"
	"fmt"
	"os"
)

type Runner struct {
	options *utils.Options
	conf    *conf.Conf
}

func NewRunner(options *utils.Options) *Runner {
	return &Runner{
		options: options,
		conf:    conf.InitConf(),
	}
}

func (r *Runner) check() (api.Apier, error) {
	var apiObj api.Apier
	switch r.options.SearchType {
	case "QuakeSearch":
		apiObj = api.NewQuake(r.conf.QuakeToken, r.options.Num)
		//case "fofa":
		//	apiObj = api.NewFofa(r.conf.FofaEmail, r.conf.FofaToken)
		//case "dnsgrep":
		//	apiObj = api.NewDsnGrep(r.conf.DnsgrepToken)
	}
	ok := apiObj.Auth()
	if ok {
		return apiObj, nil
	}
	return nil, api.NewApiError(r.options.SearchType)
}

func (r *Runner) Search() {
	apiObj, err := r.check()
	if err != nil {
		fmt.Println("Api auth err :", err)
		os.Exit(0)
	}

	result := api.Search(apiObj, r.options.SearchString)

	var UniqueResult []string
	tmpList := make(map[string]int)

	for _, item := range result {
		if item == "" {
			continue
		}
		if _, ok := tmpList[item]; !ok {
			UniqueResult = append(UniqueResult, item)
			tmpList[item] = 1
		}
	}

	//UniqueResult := []string{"http://lx1.xhu.edu.cn", "http://cmamt15.xhu.edu.cn", "http://scsz.xhu.edu.cn", "http://jwc.xhu.edu.cn", "https://panabityjdw.xhu.edu.cn", "http://face.xhu.edu.cn", "https://2022-ieeeicassp-org-s.tsgvpn.xhu.edu.cn", "https://chem-cnki-net-s.tsgvpn.xhu.edu.cn", "https://xt-cnki-net-s.tsgvpn.xhu.edu.cn", "https://202-115-153-140-8001-p.tsgvpn.xhu.edu.cn", "https://xdyy-cnki-net-s.tsgvpn.xhu.edu.cn", "https://t-go-sohu-com.tsgvpn.xhu.edu.cn", "https://data-cnki-net-s.tsgvpn.xhu.edu.cn", "https://news-xhu-edu-cn.tsgvpn.xhu.edu.cn", "https://hypt02-cnki-net-s.tsgvpn.xhu.edu.cn", "https://groupyd-chaoxing-com-s.tsgvpn.xhu.edu.cn", "https://cxjc-cnki-net-s.tsgvpn.xhu.edu.cn", "https://esi-help-clarivate-com.tsgvpn.xhu.edu.cn", "https://www-ieee--jas-net.tsgvpn.xhu.edu.cn", "http://tsgvpn.xhu.edu.cn:8118", "http://cxcyxy.xhu.edu.cn", "http://global.xhu.edu.cn", "https://face.xhu.edu.cn", "http://zzb.xhu.edu.cn", "http://jw.xhu.edu.cn", "http://nmc.xhu.edu.cn", "http://finance.xhu.edu.cn", "http://shfz.xhu.edu.cn", "http://energy.xhu.edu.cn", "http://zb.xhu.edu.cn"}
	fmt.Println("Total target url number is: ", len(UniqueResult))

	pool := utils.NewPool(UniqueResult)
	pool.Start()

}
