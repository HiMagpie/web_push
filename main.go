package main

import (
	_ "web_push/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/public", "public")
	beego.Run()
}