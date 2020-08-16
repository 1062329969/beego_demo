package routers

import (
	"beego_test/controllers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.NSGet("/test3", func(ctx *context.Context) {
		ctx.Output.Body([]byte("info"))
	})
	beego.Router("/test", &controllers.TestController{}, "get:Test_get")
	beego.Router("/test", &controllers.TestController{}, "post:Test_post")
	beego.Router("/test2", &controllers.TestController{}, "post:Test2")
	beego.Router("/any", &controllers.TestController{}, "*:Anytest")

	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/user",
				// /api/ios/create/node/
				beego.NSGet("/info", func(ctx *context.Context) {
					ctx.Output.Body([]byte("info"))
				}),
				// /api/ios/create/topic/
				beego.NSGet("/list", func(ctx *context.Context) {
					ctx.Output.Body([]byte("list"))
				}),
			),
			beego.NSGet("/admin", func(context *context.Context) {
				context.Output.Body([]byte("admin"))
			}),
		)
	beego.AddNamespace(ns)
}
