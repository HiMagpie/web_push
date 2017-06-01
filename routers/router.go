package routers

import (
	"web_push/controllers"
	"github.com/astaxie/beego"
	"web_push/controllers/center"
	"web_push/controllers/api"
	"github.com/astaxie/beego/context"
)

func init() {
	home := &controllers.HomeController{}
	beego.Router("/", home, "get:Get")
	beego.AutoRouter(home)

	beego.Router("/dashboard", &center.DashboardController{}, "get:Get")
	beego.AutoRouter(&center.DashboardController{})

	// 提供第三方的API
	apiNs := beego.NewNamespace("/api",
		beego.NSCond(func(ctx *context.Context) bool {
			return true;
		}),
		beego.NSAutoRouter(&api.AppController{}),
	)
	beego.AddNamespace(apiNs)

}
