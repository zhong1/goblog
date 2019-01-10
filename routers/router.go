package routers

import (
	"goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    ns := beego.NewNamespace("/v1",
    	beego.NSNamespace("/users",
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
    		),
    	)

    beego.AddNamespace(ns)


}
