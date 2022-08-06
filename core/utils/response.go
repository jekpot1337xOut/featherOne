package utils

type Response struct {
	ApiResults interface{}
	Entry
}

func NewResponse(apiResults interface{}, url string) *Response {
	return &Response{
		ApiResults: apiResults,
		Entry:      Entry{URL: url},
	}
}

type Entry struct {
	URL        string
	Title      string
	StatusCode int
	RawBody    []byte
	StringBody string
}
