package controllers

import (
	"github.com/astaxie/beego"
	"requestbingo/models"
	"requestbingo/storage"
	"strings"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Home() {
	recents := make([]models.Bin, 0)
	if c.GetSession("recent") != nil {
		recent, _ := c.GetSession("recent").([]string)
		for _, name := range recent {
			bin := storage.LookupBin(name)
			recents = append(recents, *bin)
		}
	}
	c.Data["recent"] = recents
	c.TplName = "home.tpl"
}

func (c *ViewController) Bin() {
	name := c.Ctx.Input.Param(":name")
	bin := storage.LookupBin(name)
	if bin == nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.Ctx.ResponseWriter.Write([]byte("Not found"))
		return
	}
	segments := strings.Split(c.Ctx.Request.RequestURI, "?")
	if len(segments) > 1 && segments[1] == "inspect" {
		if bin.Private && c.GetSession(bin.Name) != bin.SecretKey {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Ctx.ResponseWriter.Write([]byte("Private Bin"))
			return
		}

		recent, _ := c.GetSession("recent").([]string)
		if recent == nil {
			recent = []string{name}
		} else {
			remove(recent, name)
			recent = append(recent, name)
			if len(recent) > 10 {
				recent = recent[0:10]
			}
		}
		c.SetSession("recent", recent)
		c.Data["bin"] = bin
		c.Data["color"] = models.GetColor(bin)
		c.Data["base_url"] = c.Ctx.Input.Scheme() + "://" + c.Ctx.Request.Host
		c.TplName = "bin.tpl"
	} else {
		request := models.CreateRequest(c.Ctx)
		storage.CreateRequest(bin, request)
		c.Ctx.ResponseWriter.Header().Set("Sponsored-By", "https://www.runscope.com")
		c.Ctx.ResponseWriter.Write([]byte("ok"))
	}
}

func remove(slice []string, elem string) []string {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return remove(slice, elem)
			break
		}
	}
	return slice
}
