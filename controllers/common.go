package controllers

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

// func IsUserLoggedIn(c *BaseController) {
// 	GetCurrentUser(c)
// 	log.Print("-----------------")
// 	log.Print("here")
// 	log.Print("-----------------")
// }

// GetCurrentUser is ...
func GetCurrentUser(token string) {
	// token := c.Ctx.GetCookie("access_token")
	// log.Print("-----------------")
	// log.Print(token)
	// log.Print("-----------------")

}
