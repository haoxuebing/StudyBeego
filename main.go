package main

import (
	_ "StudyBeego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	//注册Orm
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qwer1234@tcp(39.104.166.99:3306)/test")
}

func main() {
	o:=orm.NewOrm()
	o.Using("default")
	//设置日志
	beego.SetLogger(logs.AdapterFile, `{"filename":"logs/test.log","daily":true}`)
	//设置允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	})

	beego.Run()
}
