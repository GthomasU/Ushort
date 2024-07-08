package api 

type ResponseBadRequest struct {
	ErrorCode    string //`json: errorCode`
	ErrorMessage string //`json:errorMessage`
}
