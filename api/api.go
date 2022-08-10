package api

import (
	"fmt"
)

type Apier interface {
	// TODO
	// dnsgrep ,fofa ,webscan's Auth return is error. need to modify.
	Auth() bool
	search(string) IPLists
}

type AuthError struct {
	name string
}

func NewApiError(name string) *AuthError {
	return &AuthError{
		name: name,
	}
}

func (a *AuthError) Error() string {
	return fmt.Sprintf("[-] %s api Authentication Error", a.name)
}

// Search entry of api
func Search(apier Apier, searchString string) IPLists {
	// TODO
	// unhandle search
	iplist := apier.search(searchString)
	return iplist
}

func SearchIp(s string) string {
	ip := webScan{}.searchIP(s)
	return ip
}

func SearchSip(s string) IPLists {
	var sameIpResult IPLists
	tmpSip := webScan{}.searchSip(s)
	sameIpResult = append(sameIpResult, tmpSip...)
	tmpSip = rapidDns{}.searchSip(s)
	sameIpResult = append(sameIpResult, tmpSip...)
	return sameIpResult
}

// SearchWeight search weight entry
func SearchWeight(s string) (int, error) {
	w, err := aizhan{}.searchWeight(s)
	if err != nil {
		return -1, nil
	}
	return w, nil
}
