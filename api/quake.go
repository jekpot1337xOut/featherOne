package api

import (
	"bytes"
	"encoding/json"
	"featherOne/Logs"
	"featherOne/core/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type IPLists []string

type Quake struct {
	Token string
	*QuakeSearchFiled
}

// NewQuake construct of Quake struct
func NewQuake(token string, num int) *Quake {
	return &Quake{
		Token:            token,
		QuakeSearchFiled: NewQuakeSearchFiled("", num),
	}
}

// Auth check the api can use or not
func (q *Quake) Auth() bool {
	INFOURL := "https://quake.360.cn/api/v3/user/info"
	request := utils.NewRequest("GET", INFOURL, nil)
	request.SetHeader("X-QuakeToken", q.Token)
	resp := request.Do()
	if !strings.Contains(resp.Entry.StringBody, "Successful") {
		return false
	}
	return true
}

// search main logic of Quake query
func (q *Quake) search(search string) IPLists {
	var iplist IPLists
	INFOURL := "https://quake.360.cn/api/v3/search/quake_service"
	q.QuakeSearchFiled.Query = search

	body, err := json.Marshal(q.QuakeSearchFiled)
	if err != nil {
		Logs.Error.Println("[-] Unmarshal quake search parameters Error")
	}
	readerBody := bytes.NewBuffer(body)
	request := utils.NewRequest("POST", INFOURL, readerBody)
	request.SetHeader("X-QuakeToken", q.Token)
	request.SetHeader("Content-Type", "application/json")
	request.UnmarshalResp(NewQuakeSearchResult())
	resp := request.Do()

	value, _ := resp.ApiResults.(*QuakeSearchResult)
	for _, item := range value.Data {
		var ip, host string
		switch item.Port {
		case 80:
			ip = fmt.Sprintf("http://%s", item.Ip)
			host = func(s string) string {
				if host != "" {
					return fmt.Sprintf("http://%s", s)
				}
				return ""
			}(item.Service.Http.Host)
		case 443:
			ip = fmt.Sprintf("https://%s", item.Ip)
			host = func(s string) string {
				if host != "" {
					return fmt.Sprintf("https://%s", s)
				}
				return ""
			}(item.Service.Http.Host)
		default:
			ip = fmt.Sprintf("http://%s:%d", item.Ip, item.Port)
			host = func(s string) string {
				if host != "" {
					return fmt.Sprintf("http://%s:%d", s, item.Port)
				}
				return ""
			}(item.Service.Http.Host)
		}
		iplist = append(iplist, ip)
		if host != "" {
			iplist = append(iplist, host)
		}
	}

	return iplist
}

// QuakeSearchFiled quake query interface parameters
type QuakeSearchFiled struct {
	Query       string      `json:"query"` // Query sentence
	Start       int         `json:"start"` // Paging start
	Size        int         `json:"size"`  // Paging Size
	IgnoreCache interface{} `json:"ignore_cache"`
	StartTime   string      `json:"start_time"` // Query start time
	EndTime     string      `json:"end_time"`   // Inquiry off time
	Include     []string    `json:"include""`   //Containing fields
}

// NewQuakeSearchFiled construct of QuakeSearchFiled struct
func NewQuakeSearchFiled(query string, size int) *QuakeSearchFiled {
	return &QuakeSearchFiled{
		Query:       query,
		Start:       0,
		Size:        size,
		IgnoreCache: false,
		StartTime:   strconv.Itoa(time.Now().Year()-1) + time.Now().Format("2006-01-02 03:04:05")[4:10] + " 00:00:00",
		EndTime:     time.Now().Format("2006-01-02 03:04:05")[:10] + " 00:00:00", // Data from the default query for the past year
		Include:     []string{"ip", "port", "service.http.host"},
	}
}

// QuakeSearchResult Quake service data interface return data structure
type QuakeSearchResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Service struct {
			Http struct {
				Host string `json:"host"`
			} `json:"http"`
		} `json:"service"`
		Port int    `json:"port"`
		Ip   string `json:"ip"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Count     int `json:"count"`
			PageIndex int `json:"page_index"`
			PageSize  int `json:"page_size"`
			Total     int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}

// NewQuakeSearchResult construct of QuakeSearchResult struct
func NewQuakeSearchResult() *QuakeSearchResult {
	return &QuakeSearchResult{}
}
