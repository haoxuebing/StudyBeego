package routers


import (
	"StudyBeego/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

func init() {
    //基础路由
    beego.Get("/base", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Post("/alice",func(ctx *context.Context){
		ctx.Output.Body([]byte("bob"))
	})
    beego.Any("/any", func(ctx *context.Context) {    //可响应任何Http请求
		ctx.Output.Body([]byte("bar"))
	})

    //正则路由
    beego.Get("/api/*.*", func(ctx *context.Context) {
		path:= ctx.Input.Param(":path")
		ext:=ctx.Input.Param(":ext")
		ctx.WriteString(path+"\n"+ext)
	})

    //RESTful 路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user",&controllers.StudentController{})

	//注解路由
	beego.Include(&controllers.CMSController{})

    //路由前置处理
    beego.InsertFilter("/",beego.BeforeRouter, func(ctx *context.Context) {
		str:=beego.AppConfig.String("appname")
		fmt.Println(str)

	})

    //自动路由 /路由名/方法名
	beego.AutoRouter(&controllers.TestUrlController{})

	beego.Router("/api2/list", &controllers.TestUrlController{}, "*:List")
	beego.Router("/person/:last/:first", &controllers.TestUrlController{})

    beego.AutoRouter(&controllers.TestOrmController{})

}
