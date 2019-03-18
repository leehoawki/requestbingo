package models

import (
	"github.com/astaxie/beego/context"
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

func CreateRequest(context *context.Context) *Request {
	request := new(Request)
	if context != nil {
		request.Id = TinyId(6)
		request.Time = time.Now()
		request.RemoteAddr = context.Request.Header.Get("X-Forwarded-For")
		if request.RemoteAddr == "" {
			request.RemoteAddr = context.Request.RemoteAddr
		}
		request.Method = context.Request.Method

		headers := make(map[string]string)
		for key, val := range context.Request.Header {
			if len(val) > 0 {
				headers[key] = val[0]
			}
		}
		request.Headers = &headers

		request.QueryString = context.Request.
		//request.FormData
		//request.Body
		request.ContentType = context.Request.Header.Get("Content-type")

		request.Body = string(context.Input.RequestBody)
		request.Path = strings.Split(context.Request.RequestURI, "?")[0]
		//request.Raw =
		request.ContentLength = context.Request.ContentLength

	}
	return request
}
