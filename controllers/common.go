package controllers

import (
	"beego-members-api/models"

	"github.com/beego/beego/v2/client/orm"
)

// DefaultSuccessResponse is ...
type DefaultSuccessResponse struct {
	// TODO: enumをサポートしたらenumに変更する
	Message string `json:"message" required:"true" example:"success" description:"result status"`
}

// DefaultErrorResponse is ...
type DefaultErrorResponse struct {
	// TODO: enumをサポートしたらenumに変更する
	Message string `json:"message" required:"true" example:"snaked_params" description:"result status"`
}

// IsUserLoggedIn is ...
func IsUserLoggedIn(userID int) bool {
	if userID == 0 {
		return false
	}
	return true
}

// GetCurrentUser is ...
func GetCurrentUser(token string) int {
	var accessToken models.AccessToken
	o := orm.NewOrm()
	err := o.QueryTable("access_token").Filter("token", token).One(&accessToken)
	if err == orm.ErrNoRows {
		return 0
	}
	return int(accessToken.User.Id)
}
