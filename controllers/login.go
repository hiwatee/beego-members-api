package controllers

import (
	"beego-members-api/models"
	"encoding/json"
	"log"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

//  UserController operations for User
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
}

type LoginRequest struct {
	Email    string `orm:"size(128)" json:"email"  example:"info@example.com"`
	Password string `orm:"size(128)" json:"password"  example:"password"`
}

type LoginResponse struct {
	Message string `json:"message" example:"success"`
}

// Login
// @Title Login
// @Description Login
// @Param   body        body    controllers.LoginRequest   true        "Login Request"
// @Success 201 {object} controllers.LoginResponse
// NOTE: FailureのBodyにObjectを渡すことが出来ないのでSuccessで代用しています
// @Success 403 {object} controllers.DefaultErrorResponse
// @router / [post]
func (c *LoginController) Login() {
	c.Ctx.Output.SetStatus(201)
	var v LoginRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		log.Fatal(err)
	}
	o := orm.NewOrm()

	// ユーザーがいるかどうか調べる
	var user models.User
	err := o.QueryTable("user").Filter("Email", v.Email).One(&user)
	if err == orm.ErrNoRows {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = DefaultErrorResponse{Message: "login_failure"}
		c.ServeJSON()
		return
	}
	// パスワード照合 | 間違ったら403
	if !user.CheckPassword(v.Password) {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = DefaultErrorResponse{Message: "login_failure"}
		c.ServeJSON()
		return
	}

	// set_cookie
	mes := DefaultSuccessResponse{Message: "success"}
	c.Data["json"] = mes
	c.ServeJSON()
}
