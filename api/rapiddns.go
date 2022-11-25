package api

import (
	"fmt"
	"regexp"
	"strings"

	"featherOne/core/utils"
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
	INFOURL := fmt.Sprintf("http://api.scrape.do/?token=ad3d00b0025842afb0f1620cf7f3301dd3ddcc23d1a&url=https://rapiddns.io/subdomain/%s?full=1#result", s)
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
