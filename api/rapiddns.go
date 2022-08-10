package api

import (
	"featherOne/core/utils"
	"fmt"
	"regexp"
	"strings"
)

// https://rapiddns.io/
type rapidDns struct {
}

// searchSip search same ip about input via rapiddns
func (rapidDns) searchSip(s string) IPLists {
	var sameSip IPLists
	INFOURL := fmt.Sprintf("https://rapiddns.io/sameip/%s?full=1#result", s)
	request := utils.NewRequest("GET", INFOURL, nil)
	resp := request.Do()

	sameSip = extractInfo(resp.StringBody)
	return sameSip
}

// searchSubdomain search subdomain via  rapiddns
func (rapidDns) searchSubdomain(s string) IPLists {
	var subdomainList IPLists
	INFOURL := fmt.Sprintf("https://rapiddns.io/subdomain/%s?full=1#result", s)
	req := utils.NewRequest("GET", INFOURL, nil)
	resp := req.Do()

	subdomainList = extractInfo(resp.StringBody)
	return subdomainList
}

// extractInfo extract needed info from html page
func extractInfo(s string) IPLists {
	var ipList IPLists
	pattern := regexp.MustCompile("</th>\n<td>(.*?)</td>")
	result := pattern.FindAllStringSubmatch(s, -1)
	for _, item := range result {
		ipList = append(ipList, strings.TrimSpace(item[1]))
	}
	return ipList
}
