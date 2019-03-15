package models

import "time"

type Color struct {
	R int
	G int
	B int
}

type Bin struct {
	Created    time.Time
	Private    bool
	Color      Color
	Name       string
	FaviconUrl string
	Requests   *[]Request
	SecretKey  string
}
