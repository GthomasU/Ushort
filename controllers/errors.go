package controllers

type ResponseBadRequest struct {
	ErrorCode    string //`json: errorCode`
	ErrorMessage string //`json:errorMessage`
}
