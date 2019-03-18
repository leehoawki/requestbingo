package models

import (
	"net/http"
    "strings"
    "time"
)

type Request struct {
	Id            string
	Time          time.Time
	RemoteAddr    string
	Method        string
	Headers       *map[string]string
	QueryString   *map[string]string
	FormData      *[]string
	Body          string
	Path          string
	ContentType   string
	Raw           string
	ContentLength int64
}

func CreateRequest(_request *http.Request) *Request {
	request := new(Request)
	if _request != nil {
		request.Id = TinyId(6)
		request.Time = time.Now()
		request.RemoteAddr = _request.Header.Get("X-Forwarded-For")
		if request.RemoteAddr == "" {
			request.RemoteAddr = _request.RemoteAddr
		}
		request.Method = _request.Method

		headers := make(map[string]string)
		for key, val := range _request.Header {
			if len(val) > 0 {
				headers[key] = val[0]
			}
		}
		request.Headers = &headers

		request.QueryString = _request.
		//request.FormData
		//request.Body
		request.ContentType = _request.Header.Get("Content-type")

		request.Body = _request.GetBody
		request.Path = strings.Split(_request.RequestURI, "?")[0]
		//request.Raw =
		request.ContentLength = _request.ContentLength

	}
	return request
}
