package main

import (
	"github.com/astaxie/beego"
	_ "github.com/leehoawki/requestbingo/models"
	_ "github.com/leehoawki/requestbingo/routers"
)

func main() {
	beego.Run()
}
