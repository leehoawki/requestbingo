package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/leehoawki/requestbingo/storage"
)

type StatController struct {
	beego.Controller
}

func (c *StatController) Stats() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	m := make(map[string]interface{})
	m["bin_count"] = len(storage.Mem.Bins)
	m["request_count"] = storage.Mem.RequestCount
	data, _ := json.Marshal(m)
	c.Ctx.ResponseWriter.Write([]byte(data))
}
