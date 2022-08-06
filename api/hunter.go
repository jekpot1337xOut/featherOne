package api

import (
	"encoding/base64"
	"featherOne/core/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Hunter struct {
	Apikey string
	HunterSearchFiled
}

func NewHunter(s string, size int) *Hunter {
	return &Hunter{
		Apikey:            s,
		HunterSearchFiled: NewHunterSearchFiled("", size),
	}
}

func (h *Hunter) Auth() bool {
	INFOURL := "https://hunter.qianxin.com/openApi/search" + "?api-key=" + h.Apikey
	request := utils.NewRequest("GET", INFOURL, nil)
	resp := request.Do()
	if !strings.Contains(resp.Entry.StringBody, "搜索内容不能为空") {
		return false
	}
	return true
}

func (h Hunter) search(s string) IPLists {
	var iplist IPLists
	h.HunterSearchFiled.search = fixHunterSearchString(s)
	INFOURL := "https://hunter.qianxin.com/openApi/search" + "?api-key=" + h.Apikey +
		hunterSearchTrans(h.HunterSearchFiled)

	request := utils.NewRequest("GET", INFOURL, nil)
	request.UnmarshalResp(NewHunterSearchResult())
	resp := request.Do()

	value, _ := resp.ApiResults.(*HunterSearchResult)
	for _, item := range value.Data.Arr {
		var ip, domain string
		switch item.Port {
		case 80:
			ip = fmt.Sprintf("http://%s", item.Ip)
			domain = func(s string) string {
				if s != "" {
					return fmt.Sprintf("http://%s", s)
				}
				return ""
			}(item.Domain)
		case 443:
			ip = fmt.Sprintf("https://%s", item.Ip)
			domain = func(s string) string {
				if s != "" {
					return fmt.Sprintf("https://%s", s)
				}
				return ""
			}(item.Domain)
		default:
			ip = fmt.Sprintf("http://%s:%d", item.Ip, item.Port)
			domain = func(s string) string {
				if s != "" {
					return fmt.Sprintf("http://%s:%d", s, item.Port)
				}
				return ""
			}(item.Domain)
		}
		iplist = append(iplist, ip)
		if domain != "" {
			iplist = append(iplist, domain)
		}
	}

	return iplist
}

type HunterSearchFiled struct {
	search     string
	page       int
	pageSize   int
	start_time string
	end_time   string
}

func NewHunterSearchFiled(search string, size int) HunterSearchFiled {
	return HunterSearchFiled{
		search:     search,
		page:       1,
		pageSize:   size,
		start_time: fmt.Sprintf("%s+00%%3A00%%3A00", strconv.Itoa(time.Now().Year()-1)+time.Now().Format("2006-01-02 03:04:05")[4:10]),
		end_time:   fmt.Sprintf("%s+23%%3A59%%3A59", time.Now().Format("2006-01-02 03:04:05")[:10]),
	}
}

func hunterSearchTrans(h HunterSearchFiled) string {
	getParameter := fmt.Sprintf(
		"&search=%s&page=%v&page_size=%v&is_web=3&start_time=%s&end_time=%s",
		base64.URLEncoding.EncodeToString([]byte(h.search)),
		h.page,
		h.pageSize,
		h.start_time,
		h.end_time,
	)
	return getParameter
}

// HunterSearchResult Hunter query data interface return data structure
type HunterSearchResult struct {
	Code int `json:"code"`
	Data struct {
		AccountType string `json:"account_type"`
		Total       int    `json:"total"`
		Time        int    `json:"time"`
		Arr         []struct {
			IsRisk         string `json:"is_risk"`
			Url            string `json:"url"`
			Ip             string `json:"ip"`
			Port           int    `json:"port"`
			WebTitle       string `json:"web_title"`
			Domain         string `json:"domain"`
			IsRiskProtocol string `json:"is_risk_protocol"`
			Protocol       string `json:"protocol"`
			BaseProtocol   string `json:"base_protocol"`
			StatusCode     int    `json:"status_code"`
			Component      []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"component"`
			Os        string `json:"os"`
			Company   string `json:"company"`
			Number    string `json:"number"`
			Country   string `json:"country"`
			Province  string `json:"province"`
			City      string `json:"city"`
			UpdatedAt string `json:"updated_at"`
			IsWeb     string `json:"is_web"`
			AsOrg     string `json:"as_org"`
			Isp       string `json:"isp"`
			Banner    string `json:"banner"`
		} `json:"arr"`
		ConsumeQuota string `json:"consume_quota"`
		RestQuota    string `json:"rest_quota"`
		SyntaxPrompt string `json:"syntax_prompt"`
	} `json:"data"`
	Message string `json:"message"`
}

// NewHunterSearchResult construct of HunterSearchResult struct
func NewHunterSearchResult() *HunterSearchResult {
	return &HunterSearchResult{}
}

// fixHunterSearchString fix search string
func fixHunterSearchString(s string) string {
	splitSlice := strings.Split("domain=xhu.edu.cn", "=")
	fixString := fmt.Sprintf(`%s="%s"`, splitSlice[0], splitSlice[1])
	return fixString
}
