package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "myFirstBeego/routers"
	"github.com/astaxie/beego"
)

func init()  {
	fmt.Println("我是master分支中的init方法")
	fmt.Println("master中增加打印信息")
}
func hotfix()  {
	fmt.Println("我是hotfix分支中的hotfix函数")
	fmt.Println("master中增加打印信息")
}
func main() {
	logs.Info("\r\n 1111 from main.go, init Methods test in github")
	fmt.Println("I'm is hotfix  changing ")
	fmt.Println("I'm is  changing ")
	beego.Run()
}

