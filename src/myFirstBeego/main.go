package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "myFirstBeego/routers"
	"github.com/astaxie/beego"
)

func init()  {
	fmt.Println("我是hotfix分支中的init函数")
}
func main() {
	logs.Info("\r\n 1111 from main.go, init Methods test in github")
	fmt.Println("I'm is hotfix  changing ")
	fmt.Println("I'm is  changing ")
	beego.Run()
}

