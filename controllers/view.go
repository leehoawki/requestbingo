package controllers

import (
	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Home() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.tpl"
}

func (c *ViewController) Docs() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "doc.tpl"
}

func (c *ViewController) Bin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "bin.tpl"
}
