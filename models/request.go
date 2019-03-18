package models

import (
	"time"
)

type Request struct {
	Id            string
	Time          time.Time
	RemoteAddr    string
	Method        string
	Headers       map[string]string
	QueryString   map[string]string
	FormData      *[]string
	Body          string
	Path          string
	ContentType   string
	Raw           string
	ContentLength int
}
