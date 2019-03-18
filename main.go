package main

import (
	"github.com/astaxie/beego"
	_ "requestbingo/models"
	_ "requestbingo/routers"
)

func main() {
	beego.Run()
}
