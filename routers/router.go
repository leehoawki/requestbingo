package routers

import (
	"github.com/astaxie/beego"
	"requestbingo/controllers"
)

func init() {
	beego.Router("/", &controllers.ViewController{}, "*:Home")
	beego.Router("/:name", &controllers.ViewController{}, "*:Bin")

	beego.Router("/api/v1/bins", &controllers.ApiController{}, "post:Bins")
	beego.Router("/api/v1/bins/:name", &controllers.ApiController{}, "get:Bin")
	beego.Router("/api/v1/bins/:bin/requests", &controllers.ApiController{}, "get:Requests")
	beego.Router("/api/v1/bins/:bin/requests/:request", &controllers.ApiController{}, "get:Request")

	beego.Router("/api/v1/stats", &controllers.StatController{}, "get:Stats")
}
