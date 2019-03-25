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
	var bin *models.Bin
	if private == "true" || private == "on" {
		bin = models.CreateBin(true)
		c.SetSession(bin.Name, bin.SecretKey)
	} else {
		bin = models.CreateBin(false)
	}
	storage.CreateBin(bin)
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Write([]byte(models.Bin2JsonString(bin)))
}

func (c *ApiController) Bin() {
	name := c.Ctx.Input.Param(":name")
	bin := storage.LookupBin(name)
	if bin == nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.Ctx.ResponseWriter.Write([]byte("Not found"))
	} else {
		c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
		c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		c.Ctx.ResponseWriter.Write([]byte(models.Bin2JsonString(bin)))
	}
}

func (c *ApiController) Requests() {
	bin_name := c.Ctx.Input.Param(":bin")
	bin := storage.LookupBin(bin_name)
	if bin == nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.Ctx.ResponseWriter.Write([]byte("Bin Not found"))
	} else {
		c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
		c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		c.Ctx.ResponseWriter.Write([]byte(models.Requests2JsonString(&bin.Requests)))
	}
}

func (c *ApiController) Request() {
	bin_name := c.Ctx.Input.Param(":bin")
	request_id := c.Ctx.Input.Param(":request")
	bin := storage.LookupBin(bin_name)
	if bin == nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.Ctx.ResponseWriter.Write([]byte("Bin Not found"))
	} else {
		for _, request := range bin.Requests {
			if request.Id == request_id {
				c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
				c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
				c.Ctx.ResponseWriter.Write([]byte(models.Request2JsonString(&request)))
				return
			}
		}
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.Ctx.ResponseWriter.Write([]byte("Request Not found"))
	}
}
