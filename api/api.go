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

// Auth TODO
// complete function
func Auth(searchType string) (Apier, error) {
	return nil, nil
}
