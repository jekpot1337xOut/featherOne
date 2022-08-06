package main

import (
	"featherOne/core"
	"featherOne/core/utils"
)

func main() {
	options := utils.ParseOptions()
	r := core.NewRunner(options)
	r.Search()
}
