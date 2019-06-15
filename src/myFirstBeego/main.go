package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "myFirstBeego/routers"
)

func main() {
	logs.Info("\r\n 1111 from main.go, init Methods test in github")
	beego.Run()
	fmt.Println("test branch")
}

func homeTest()  {
	fmt.Println("在homeTest中开发项目.")
}
