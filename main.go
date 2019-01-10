package main

import (
	_ "goblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)

	beego.Run()
}

