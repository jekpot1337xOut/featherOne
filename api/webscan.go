package api

import (
	"featherOne/core/utils"
	"fmt"
)

// https://www.webscan.cc/api/
type webScan struct {
}

// searchIp seach the ip address aoubt input via webscan
func (webScan) searchIP(s string) string {
	var ip string
	req := utils.NewRequest("GET", fmt.Sprintf("https://api.webscan.cc/?action=getip&domain=%s", s), nil)
	req.UnmarshalResp(&wsIpResult{})
	resp := req.Do()

	value, _ := resp.ApiResults.(*wsIpResult)
	ip = value.Ip

	return ip
}

func (webScan) searchSip(s string) IPLists {
	var sameIpList IPLists
	req := utils.NewRequest("GET", fmt.Sprintf("http://api.webscan.cc/?action=query&ip=%s", s), nil)
	req.UnmarshalResp(&[]wsSipResult{})
	resp := req.Do()

	value, _ := resp.ApiResults.(*[]wsSipResult)
	for _, item := range *value {
		sameIpList = append(sameIpList, item.Domain)
	}

	return sameIpList
}

// wsSipResult Same IP Address query api result
// http://api.webscan.cc/?action=query&ip=www.webscan.cc
type wsSipResult struct {
	Domain string `json:"domain"`
	Title  string `json:"title"`
}

// wsIpResult Query Ip Address api result
// https://api.webscan.cc/?action=getip&domain=www.google.com
type wsIpResult struct {
	Ip   string `json:"ip"`
	Info string `json:"info"`
}
