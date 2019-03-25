package models

import (
	"encoding/json"
	"fmt"
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
	QueryString   string
	Querys        *map[string]string
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

		segments := strings.Split(context.Request.RequestURI, "?")
		request.Path = segments[0]
		querys := make(map[string]string)
		if len(segments) > 1 {
			request.QueryString = segments[1]
			params := strings.Split(context.Request.RequestURI, "?")[1]
			for _, elem := range strings.Split(params, "&") {
				kv := strings.Split(elem, "=")
				k := kv[0]
				v := kv[1]
				querys[k] = v
			}
		}
		request.Querys = &querys
		formdata := make(map[string]string)
		for key, val := range context.Request.Form {
			if len(val) > 0 {
				formdata[key] = val[0]
			}
		}
		request.FormData = &formdata

		request.ContentType = context.Request.Header.Get("Content-type")

		request.Body = string(context.Input.RequestBody)
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

func Requests2JsonString(request *[]Request) string {
	a := make([]map[string]interface{}, len(*request))

	for index, request := range *request {
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
		a[index] = m
	}
	fmt.Println(len(*request))
	data, _ := json.Marshal(a)
	return string(data)
}
