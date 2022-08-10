package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/projectdiscovery/gologger"
)

const (
	RED   = "\033[1;31;40m"
	GREEN = "\033[1;32;40m"
	END   = "\033[0m"
)

var KEYWORDS = []string{"系统", "管理", "登录", "后台", "login"}

type Analyzer interface {
	ParseResp(response *Response) *Response
}

type Analyze struct {
}

// ParseResp This function contains all the logic of all processing response
// Parse response from *http.reponse
// In this case,we store response extract title and unmarshal api json result
// You can rewrite this function to define you own logic
func ParseResp(r *Response, response *http.Response) *Response {
	defer func() {
		if recover() != nil {
			gologger.Error().Msg("search grammar err.Please check and run again")
			os.Exit(0)
		}
	}()

	rawResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		gologger.Error().Msg("Read raw body content Error")
		return nil
	}
	r.Entry.RawBody = rawResponse
	r.Entry.StringBody = string(rawResponse)
	r.Entry.StatusCode = response.StatusCode
	r.Entry.Title = getTitle(r.Entry.StringBody)

	if r.ApiResults != nil {
		err = json.Unmarshal(r.Entry.RawBody, r.ApiResults)
		if err != nil {
			r.ApiResults = nil
			gologger.Error().Msg("Unmarshal api result Error")
		}
	}
	return r
}

// getTitle extract Title from response body
func getTitle(s string) string {
	reg := regexp.MustCompile(`<title>(.*?)</title>`)
	t := reg.FindStringSubmatch(s)
	if len(t) > 0 {
		return t[1]
	}
	return ""
}

// fixUrl fix url if it's a domain format
func fixUrl(s string) string {
	if !strings.HasPrefix(s, "http") {
		return fmt.Sprintf("http://%s", s)
	}
	return s
}

// colorOut  output with color
func colorOut(resp *Response) {
	if resp.StatusCode == 200 {
		lowerTitle := strings.ToLower(resp.Title)
		for _, keyword := range KEYWORDS {
			if strings.Contains(lowerTitle, keyword) {
				fmt.Printf("%s[+] URL: %20s  StatusCode: %d %sTitle: %s %s\n", GREEN, resp.URL, resp.StatusCode, RED, resp.Title, END)
			}
		}
	}
}
