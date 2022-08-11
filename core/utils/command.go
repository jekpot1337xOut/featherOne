package utils

import (
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

type Options struct {
	FofaSearch    bool   // fofa query string
	QuakeSearch   bool   // quake query string
	HunterSearch  bool   // hunter query string
	DnsgrepSearch string // dnsgrep query string
	SearchString  string // search grammar
	AutoGrammar   bool
	Url           string // target url
	Filename      string // target url with file
	IpAddress     string // target ip address
	Num           int    // query data number
	TreadNum      int    // pool thread number

	Ip     bool // search domain's ip switch
	SameIp bool // search same ip of the input switch
	Weight bool // search weight of the domain

	Subdomain string // integrate subdomain module

	Silent bool // show more info or only subdomain
	Probe  bool
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
    \/_/\/____/\/__/\/_/ \/__/ \/_/\/_/\/____/ \/_/   \/_____/\/_/\/_/\/____/ v0.9.0
`

	gologger.Print().Msgf("%s\n", GREEN+banner+END)
}
func ParseOptions() *Options {
	options := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("A simple tool to ez information search")

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&options.Url, "url", "u", "", "target Url"),
		//flagSet.StringVarP(&options.Filename, "file", "f", "", "target url with file"), // not ready
		//flagSet.IntVarP(&options.TreadNum, "thread", "t", 15, "thread number"),
	)

	flagSet.CreateGroup("searchEngine", "SearchEngine",
		flagSet.StringVarP(&options.SearchString, "searchgrammar", "sg", "", "corresponding search engine's search grammar"),
		flagSet.BoolVarP(&options.FofaSearch, "fofa", "sfo", false, "Use fofa api to search"),
		flagSet.BoolVarP(&options.QuakeSearch, "quake", "squ", false, "Use quake api to search"),
		flagSet.BoolVarP(&options.HunterSearch, "hunter", "shu", false, "Use hunter api to search"),
		flagSet.IntVarP(&options.Num, "number", "num", 30, "Query data quantity "),
		flagSet.BoolVarP(&options.AutoGrammar, "autogrammar", "autog", false, "parse search into corresponding engine grammar"),
	)

	flagSet.CreateGroup("ipinfo", "Ipinfo",
		flagSet.BoolVar(&options.Ip, "ip", false, "search domain' ip"),
		flagSet.BoolVarP(&options.SameIp, "sameip", "sip", false, "search same ip"),
		flagSet.BoolVarP(&options.Weight, "weight", "wgt", false, "search weight"),
	)

	//flagSet.CreateGroup("integrate", "Integrate",
	//	flagSet.StringVar(&options.Subdomain, "subdomain", "", "Built-in integration subdomain module"),
	//)

	flagSet.CreateGroup("mode", "Mode",
		flagSet.BoolVar(&options.Silent, "silent", false, "Show subdomain or ip only"),
		flagSet.BoolVar(&options.Probe, "probe", false, "simple probe to get statusCode and title"),
	)

	flagSet.Parse()

	options.setOutMode()

	if !options.Silent {
		showBanner()
	}

	return options
}

func (o *Options) setOutMode() {
	if o.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}
}
