package main

import (
	"featherOne/core"
	"featherOne/core/utils"
	"github.com/projectdiscovery/gologger"
	"time"
)

func main() {
	start := time.Now()
	options := utils.ParseOptions()
	r := core.NewRunner(options)
	r.Search()
	gologger.Info().Msgf("Task done,cost: %v\n", time.Since(start))
}
