package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	fmt.Println(beego.BConfig.AppName + ":" + strconv.Itoa(beego.BConfig.Listen.HTTPPort))
	c.TplName = "index.tpl"
}
