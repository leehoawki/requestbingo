package models

import (
	"math/rand"
	"strconv"
	"time"
)

type Color struct {
	R int
	G int
	B int
}

type Bin struct {
	Created    time.Time
	Private    bool
	Color      *Color
	Name       string
	FaviconUrl string
	Requests   []Request
	SecretKey  string
}

func CreateBin(p bool) *Bin {
	bin := new(Bin)
	bin.Created = time.Now()
	bin.Private = p
	bin.Color = RandomColor()
	bin.Name = TinyId(8)
	bin.FaviconUrl = Solid16x16gifDatauri(bin.Color)
	bin.Requests = make([]Request, 0)
	if p {
		bin.SecretKey = strconv.Itoa(RandomString())
	}
	return bin
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomByte() int {
	return (5 + rand.Int()%20) * 10
}

func RandomString() int {
	return rand.Int()
}

func RandomColor() *Color {
	color := new(Color)
	color.R = RandomByte()
	color.G = RandomByte()
	color.B = RandomByte()
	return color
}
