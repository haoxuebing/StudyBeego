package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
)

type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("AllBlock", c.AllBlock)
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("aaa", c.AAA)
}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {
	this.Ctx.WriteString("this is AllBlock")
}

// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {
	this.Ctx.WriteString("this is StaticBlock")
}

// @router /aaa/ [get]
func (this *CMSController) AAA() {
	v := this.GetSession("asta")
	if v != nil {
		this.SetSession("asta",v.(int)+1)
	} else {
		this.SetSession("asta", 1)
	}
	this.Ctx.WriteString(strconv.Itoa(this.GetSession("asta").(int)))

	var v1 interface{}="\nasdf"
	this.Ctx.WriteString(strconv.Itoa(this.GetSession("asta").(int)))
	this.Ctx.WriteString(v1.(string))

	var v2 interface{}=123
	fmt.Println(v2.(int)+1)
}
