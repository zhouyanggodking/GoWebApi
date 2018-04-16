package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["GoWebApi/controllers:MainController"] = append(beego.GlobalControllerRouter["GoWebApi/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/rpc`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
