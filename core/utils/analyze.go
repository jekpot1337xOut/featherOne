package utils

import (
	"encoding/json"
	"featherOne/Logs"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
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
			fmt.Println("search grammar err.Please check and run again")
			os.Exit(0)
		}
	}()

	rawResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logs.Error.Println("[-] Read raw body content Error")
		return nil
	}
	r.Entry.RawBody = rawResponse
	r.Entry.StringBody = string(rawResponse)
	r.Entry.StatusCode = response.StatusCode
	r.Entry.Title = getTitle(r.Entry.StringBody)

	if r.ApiResults != nil {
		err = json.Unmarshal(r.Entry.RawBody, r.ApiResults)
		if err != nil {
			Logs.Error.Println("[-] Unmarshal api result Error")
			r.ApiResults = nil
		}
	}
	return r
}

// getTitle extract Title from response body
func getTitle(s string) string {
	reg := regexp.MustCompile(`<title>(.*?)</title>`)
	if reg == nil {
		fmt.Println("Regexp to extract Title errored :", reg)
		return ""
	}
	title := reg.FindString(s)
	//return title[8 : len(title)-8]
	return title
}

func colorOut(resp *Response) {
	fmt.Println(resp.URL, resp.Title)
	if resp.StatusCode == 200 {
		lowerTitle := strings.ToLower(resp.Title)
		for _, keyword := range KEYWORDS {
			if strings.Contains(lowerTitle, keyword) {
				fmt.Printf("%s[+] URL: %20s  StatusCode: %d %sTitle: %s %s\n", GREEN, resp.URL, resp.StatusCode, RED, resp.Title, END)
			}
		}
	}
}
