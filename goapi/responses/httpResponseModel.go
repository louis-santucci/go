package responses

import "net/http"

type OKResponse struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status uint        `json:"status"`
	Error  interface{} `json:"error"`
}

type JWTResponse struct {
	Status uint   `json:"status"`
	Token  string `json:"jwt_token"`
}

func NewErrorResponse(status uint, error interface{}) ErrorResponse {
	return ErrorResponse{
		Status: status,
		Error:  error,
	}
}

func NewOKResponse(data interface{}) OKResponse {
	return OKResponse{
		Status: http.StatusOK,
		Data:   data,
	}
}

func NewJWTResponse(status uint, token string) JWTResponse {
	return JWTResponse{
		Status: status,
		Token:  token,
	}
}
