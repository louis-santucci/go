package responses

import "net/http"

type OKResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

type JWTResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
	Email  string `json:"email"`
}

func NewErrorResponse(status int, error interface{}) ErrorResponse {
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

func NewJWTResponse(status int, token string, email string) JWTResponse {
	return JWTResponse{
		Status: status,
		Token:  token,
		Email:  email,
	}
}
