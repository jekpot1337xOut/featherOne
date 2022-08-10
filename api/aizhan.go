package api

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	
	"featherOne/core/utils"

	"github.com/projectdiscovery/gologger"
)

// https://baidurank.aizhan.com
type aizhan struct {
}

// searchWeight search weight via aizhan
func (aizhan) searchWeight(s string) (int, error) {
	INFOURL := fmt.Sprintf("https://baidurank.aizhan.com/baidu/%s/", s)
	req := utils.NewRequest("GET", INFOURL, nil)
	req.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Appl"+
		"eWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.SetHeader("Host", "baidurank.aizhan.com")
	req.SetHeader("Pragma", "no-cache")
	resq := req.Do()

	w, err := getWeight(resq.StringBody)

	if err != nil {
		gologger.Error().Msgf("%s\n", err)
		return -1, err
	}

	return w, nil
}

// getWeight extract weight from html
func getWeight(s string) (int, error) {
	pattern := regexp.MustCompile("images/br/([0-9])\\.png")
	w := pattern.FindStringSubmatch(s)
	if len(w) == 2 {
		weightS := w[1]
		weightI, _ := strconv.Atoi(weightS)
		return weightI, nil
	}
	return -1, errors.New("get weight failed")
}
