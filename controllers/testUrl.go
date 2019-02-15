package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TestUrlController struct {
	beego.Controller
}

func (this *TestUrlController) Get() {
	this.Data["Username"] = "astaxie"
	this.Ctx.Output.Body([]byte("get"))
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":last")))
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":frist")))

}

func (this *TestUrlController) Post() {
	this.Data["Username"] = "astaxie"
	fmt.Println(this.Ctx.Input.Param(":last"))
	fmt.Println(this.Ctx.Input.Param(":first"))
	this.Ctx.Output.Body([]byte("post"))
	//this.Ctx.Output.Body([]byte("post2"))
}

func (this *TestUrlController)Log(){
	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")

	this.Ctx.WriteString("This is beego Log")
}


func (this *TestUrlController) List() {
	this.Ctx.Output.Body([]byte("i am list"))
}

func (this *TestUrlController) Params() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Params()["0"] + this.Ctx.Input.Params()["1"] + this.Ctx.Input.Params()["2"]))
}

func (this *TestUrlController) Myext() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
}

func (this *TestUrlController) GetUrl() {
	this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))
}
