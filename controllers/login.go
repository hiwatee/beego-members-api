package controllers

import (
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
	Email   string `json:"email" required:"true" example:"info@example.com"`
	Pasword string `json:"password" required:"true" example:"password"`
}

type LoginResponse struct {
	// TODO: enumをサポートしたらenumに変更する
	Message string `json:"message" required:"true" example:"success" description:"result status"`
}

// NOTE: exampleの出し分けが出来ないので別にしています。
type LoginFailureResponse struct {
	Message string `json:"message" required:"true" example:"failed" description:"result status"`
}

// Login
// @Title Login
// @Description Login
// @Param   body        body    controllers.LoginRequest   true        "Login Request"
// @Success 201 {object} controllers.LoginResponse
// NOTE: FailureのBodyに構造体を渡すことが出来ないのでSuccessで代用しています
// @Success 403 {object} controllers.LoginFailureResponse
// @router / [post]
func (c *LoginController) Login() {
	c.Ctx.Output.SetStatus(201)
	mes := LoginResponse{Message: "success"}
	c.Data["json"] = mes
	c.ServeJSON()
}
