// @APIVersion 1.0.0
// @Title flowtest API
// @Description flow has every tool to get any job done, so codename for the new flowtest APIs.
// @Contact 504284@qq.com
package routers

import (
	"github.com/3xxx/flowtest/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//运行跨域请求
	//在http请求的响应流头部加上如下信息
	//rw.Header().Set("Access-Control-Allow-Origin", "*")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//自动化文档
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/admin",
				beego.NSInclude(
					&controllers.MainController{},
					// &controllers.CustomerCookieCheckerController{},
				),
			),
			// beego.NSNamespace("/cms",
			// 	beego.NSInclude(
			// 		&controllers.CMSController{},
			// 	),
			// ),
			// beego.NSNamespace("/suggest",
			// 	beego.NSInclude(
			// 		&controllers.SearchController{},
			// 	),
			// ),
		)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})
	beego.SetStaticPath("/swagger", "swagger")
}
