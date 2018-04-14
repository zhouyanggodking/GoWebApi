// @APIVersion 1.0.0
// @Title Godking Restapi for Testing
// @Description version 1 for testing
// @Contact zhou_jing_king@hotmail.com
package routers

import (
	"GoWebApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/api/test", &controllers.MainController{})
	// beego.Include(&controllers.MainController{})
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/person",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
