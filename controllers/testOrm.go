package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Users struct {
	Id int64
	Name string
	Gender int
	Age int
}

type TestOrmController struct {
	beego.Controller
}
func init(){
	orm.RegisterModel(new(Users))
}

func (this *TestOrmController)GetAllUser(){
	o:=orm.NewOrm()
	qs:=o.QueryTable(new(Users))
	var l []Users
	qs.All(&l)
	fmt.Println(l)
	this.Data["json"]=l
	this.ServeJSON()
	//this.Ctx.WriteString("this is orm")
}

func (this *TestOrmController)GetUsers(){
	o:=orm.NewOrm()
	var user Users
	var users []Users
	o.Raw("SELECT * FROM users where id=1").QueryRow(&user)
	num,_:= o.Raw("SELECT * FROM users").QueryRows(&users)
	fmt.Println(users)
	fmt.Println("num:",num)
	this.Data["json"]=user
	this.ServeJSON()

}