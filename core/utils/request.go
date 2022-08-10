package utils

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"featherOne/Logs"
)

type Request struct {
	*http.Request
	URL        string
	ApiResults interface{}
}

func NewRequest(method, url string, body io.Reader) *Request {
	var request *http.Request
	if method == "GET" {
		request, _ = http.NewRequest("GET", url, nil)
	} else {
		request, _ = http.NewRequest("POST", url, body)
	}

	return &Request{
		Request: request,
		URL:     url,
	}

}

// Do a Customized request and return a *Reponse
func (r *Request) Do() *Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}

	tmpResp, err := client.Do(r.Request)
	if err != nil {
		Logs.Error.Println(err)
	}
	defer tmpResp.Body.Close()

	resp := NewResponse(r.ApiResults, r.URL)
	ParseResp(resp, tmpResp)
	return resp
}

// SetHeader quickly set header
func (r *Request) SetHeader(fieldName, value string) *Request {
	r.Request.Header.Set(fieldName, value)
	return r
}

// UnmarshalResp quickly unmarshall json restful api result
// In this case, apiJsonStruct Must be pointer type
func (r *Request) UnmarshalResp(apiJsonStructPoint interface{}) *Request {
	r.ApiResults = apiJsonStructPoint
	return r
}
