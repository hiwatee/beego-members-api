// @APIVersion 1.0.0
// @Title UserMemberApi
// @Description The sample API for Beego
// @Contact tochika.biz@gmail.com
// @TermsOfServiceUrl https://github.com/hiwatee
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-members-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
