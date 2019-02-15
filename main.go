package main

import (
	_ "StudyBeego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.SetLogger(logs.AdapterFile, `{"filename":"logs/test.log","daily":true}`)

	//设置允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	})

	beego.Run()
}
