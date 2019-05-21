package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/bigdata_api/controllers"
	"github.com/bigdata_api/models"
	_ "github.com/bigdata_api/routers"
	"time"
)

func init() {
	models.Init()
	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	orm.DefaultTimeLoc = time.UTC
	beego.ErrorController(&controllers.ErrorController{})
	beego.BConfig.ServerName = "snail server 1.0"
	beego.Run()
}
