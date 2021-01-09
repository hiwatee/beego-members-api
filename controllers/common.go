package controllers

import "log"

type DefaultSuccessResponse struct {
	// TODO: enumをサポートしたらenumに変更する
	Message string `json:"message" required:"true" example:"success" description:"result status"`
}

type DefaultErrorResponse struct {
	// TODO: enumをサポートしたらenumに変更する
	Message string `json:"message" required:"true" example:"snaked_params" description:"result status"`
}

func IsUserLoggedIn() {
	log.Print("-----------------")
	log.Print("here")
	log.Print("-----------------")
}

func GetCurrentUser() {

}
