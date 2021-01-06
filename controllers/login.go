package controllers

import (
	"encoding/json"
	"fmt"
	"log"

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
	var v LoginRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id: %v", v.Email)
	// eメールでユーザーを検索 | なかったら404

	// パスワード照合 | 間違ったら403

	// set_cookie
	mes := DefaultSuccessResponse{Message: "success"}
	c.Data["json"] = mes
	c.ServeJSON()
}
