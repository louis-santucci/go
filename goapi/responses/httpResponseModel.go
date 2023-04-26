package responses

import (
	"louissantucci/goapi/models"
	"net/http"
)

type OKResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

type JWTResponse struct {
	Status int              `json:"status"`
	Token  string           `json:"token"`
	User   *models.UserInfo `json:"user"`
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

func NewJWTResponse(status int, token string, userInfo *models.UserInfo) JWTResponse {
	return JWTResponse{
		Status: status,
		Token:  token,
		User:   userInfo,
	}
}
