package main

import (
	"github.com/astaxie/beego/logs"
	_ "myFirstBeego/routers"
	"github.com/astaxie/beego"
)

func main() {
	logs.Info("\r\n hello world from main.go, init Methods test in github")
	beego.Run()
}

