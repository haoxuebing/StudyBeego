package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type StudentController struct {
	beego.Controller
}

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func (c *StudentController) Get() {
	//c.Data["Website"] = "hello world"
	//c.Data["Email"] = "user@gmail.com"
	//c.TplName = "index.tpl"
	id := c.GetString("id", "123")
	//c.Ctx.WriteString(id)
	//c.Data["id"]=id
	s := new(Student)
	s.Name=id
	s.Age=12
	s.Price=12.00
	s.Classes=[]string{"asdf","123"}
	b,_:=json.Marshal(s)
	//c.Ctx.WriteString(string(b))
	c.Data["json"]=string(b)
	c.ServeJSON()

}
