package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) PageUser() {
	logs.Info("\r\n from router.go init Methods")
	u.TplName = "admin.html"
}
