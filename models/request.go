package models

import (
	"encoding/json"
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
	FormData      *map[string]string
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

		params := context.Input.Params()
		request.QueryString = &params

		formdata := make(map[string]string)
		for key, val := range context.Request.Form {
			if len(val) > 0 {
				formdata[key] = val[0]
			}
		}
		request.FormData = &formdata

		request.ContentType = context.Request.Header.Get("Content-type")

		request.Body = string(context.Input.RequestBody)
		request.Path = strings.Split(context.Request.RequestURI, "?")[0]
		request.Raw = string(context.Input.RequestBody)
		request.ContentLength = context.Request.ContentLength
	}
	return request
}

func Request2JsonString(request *Request) string {
	m := make(map[string]interface{})
	m["id"] = request.Id
	m["time"] = request.Time
	m["remote_addr"] = request.RemoteAddr
	m["method"] = request.Method
	m["headers"] = request.Headers
	m["query_string"] = request.QueryString
	m["raw"] = request.Raw
	m["form_data"] = request.FormData
	m["body"] = request.Body
	m["path"] = request.Path
	m["content_length"] = request.ContentLength
	m["content_type"] = request.ContentType
	data, _ := json.Marshal(m)
	return string(data)
}
