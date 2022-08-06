package utils

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
)

type Options struct {
	FofaSearch    string // fofa query string
	QuakeSearch   string // quake query string
	DnsgrepSearch string // dnsgrep query string
	HunterSearch  string // hunter query string
	SearchString  string
	Url           string // target url
	Filename      string // target url with file
	IpAddress     string // target ip address
	SearchType    string // which cyberspace you choose to use
	Num           int    // query data number
	TreadNum      int    // pool thread number

}

func showBanner() {
	// http://www.network-science.de/ascii/  larry3d
	banner := `
 ____                 __    __                     _____                     
/\  _>\              /\ \__/\ \                   /\  __>\                   
\ \ \L\_\ __     __  \ \ ,_\ \ \___      __   _ __\ \ \/\ \    ___      __   
 \ \  _\/'__>\ /'__>\ \ \ \/\ \  _ >\  /'__>\/\>'__\ \ \ \ \ /' _ >\  /'__>\ 
  \ \ \/\  __//\ \L\.\_\ \ \_\ \ \ \ \/\  __/\ \ \/ \ \ \_\ \/\ \/\ \/\  __/ 
   \ \_\ \____\ \__/.\_\\ \__\\ \_\ \_\ \____\\ \_\  \ \_____\ \_\ \_\ \____\
    \/_/\/____/\/__/\/_/ \/__/ \/_/\/_/\/____/ \/_/   \/_____/\/_/\/_/\/____/
`

	fmt.Println(GREEN + banner + END)
}
func ParseOptions() *Options {
	options := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("A simple tool to ez information search")

	flagSet.CreateGroup("input", "Input mode(url/file)",
		flagSet.StringVarP(&options.Url, "url", "u", "", "target Url"),
		flagSet.StringVarP(&options.Filename, "file", "f", "", "target url with file"),
		flagSet.IntVarP(&options.TreadNum, "thread", "t", 15, "thread number"),
	)

	flagSet.CreateGroup("search", "search with Cyberspace search Engines apis(quake/fofa/dnsgrep)",
		flagSet.StringVarP(&options.FofaSearch, "fofasearch", "sfo", "", "Use fofa api to search"),
		flagSet.StringVarP(&options.QuakeSearch, "quakesearch", "squ", "", "Use quake api to search"),
		flagSet.StringVarP(&options.DnsgrepSearch, "dnsgrepsearch", "sdg", "", "Use dnsgrep api to search"),
		flagSet.StringVarP(&options.HunterSearch, "huntersearch", "shu", "", "Use hunter api to search"),
		flagSet.IntVarP(&options.Num, "number", "num", 30, "Query data quantity "),
	)

	// TODO
	// Add ip2domain api
	// Add domain2ip api
	flagSet.CreateGroup("search weight", "search target website' weight")
	flagSet.CreateGroup("ip info", "search ip information")

	flagSet.CreateGroup("Output", "output")

	showBanner()

	flagSet.Parse()

	if options.FofaSearch != "" {
		options.SearchType = "FofaSearch"
		options.SearchString = options.FofaSearch
	} else if options.QuakeSearch != "" {
		options.SearchType = "QuakeSearch"
		options.SearchString = options.QuakeSearch
	} else if options.DnsgrepSearch != "" {
		options.SearchType = "DnsgrepSearch"
		options.SearchString = options.DnsgrepSearch
	} else if options.HunterSearch != "" {
		options.SearchType = "HunterSearch"
		options.SearchString = options.HunterSearch
	}

	return options
}
