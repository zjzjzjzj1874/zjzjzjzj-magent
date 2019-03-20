package routers

import (
	"github.com/astaxie/beego/logs"
	"myFirstBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	logs.Info("\r\n from router.go init Methods")
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/user", &controllers.UserController{})
    beego.Router("/user", &controllers.UserController{},"get:PageUser")
	beego.SetStaticPath("/static","public")
}
