package routers

import (
	"github.com/astaxie/beego"
	"requestbingo/controllers"
)

func init() {
	beego.Router("/", &controllers.ViewController{}, "*:Home")
	beego.Router("/:path", &controllers.ViewController{}, "*:Bin")
	beego.Router("/docs/:name", &controllers.ViewController{}, "*:Docs")

	beego.Router("/api/v1/bins", &controllers.ViewController{})
	beego.Router("/api/v1/bins/:name", &controllers.ViewController{})
	beego.Router("/api/v1/bins/:bin/requests", &controllers.ViewController{})
	beego.Router("/api/v1/bins/:bin/requests/:name", &controllers.ViewController{})

	beego.Router("/api/v1/stats", &controllers.ViewController{})
}
