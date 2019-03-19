package controllers

import (
	"github.com/astaxie/beego"
	"requestbingo/models"
	"requestbingo/storage"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Bins() {
	private := c.Ctx.Request.Form.Get("private")
	if private == "true" || private == "on" {
		bin := models.CreateBin(true)
		c.SetSession(bin.Name, bin.SecretKey)
	} else {
		bin := models.CreateBin(false)
	}

}

func (c *ApiController) Bin() {
	name := c.GetString("name")
	storage.LookupBin(name)
}

func (c *ApiController) Requests() {

}

func (c *ApiController) Request() {

}

func response(code int) {

}
