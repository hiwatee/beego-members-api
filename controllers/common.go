package controllers

type DefaultSuccessResponse struct {
	// TODO: enumをサポートしたらenumに変更する
	Message string `json:"message" required:"true" example:"success" description:"result status"`
}
