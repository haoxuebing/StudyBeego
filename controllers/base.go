package controllers

import (
	"github.com/astaxie/beego"
	"mytest/models"
	"time"
)

type NestPreparer interface {
	NestPrepare()
}

// baseRouter implemented global settings for all other routers.
type BaseController struct {
	beego.Controller
	//i18n.Locale
	user    models.User
	isLogin bool
}
// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {

	// page start time
	this.Data["PageStartTime"] = time.Now()

	// Setting properties.
	//this.Data["AppDescription"] = utils.AppDescription
	//this.Data["AppKeywords"] = utils.AppKeywords
	this.Data["AppName"] = beego.Config{}.AppName
	//this.Data["AppVer"] = utils.AppVer
	//this.Data["AppUrl"] = utils.AppUrl
	//this.Data["AppLogo"] = utils.AppLogo
	//this.Data["AvatarURL"] = utils.AvatarURL
	//this.Data["IsProMode"] = utils.IsProMode

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}