package models

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Token   string `json:"jwt_token"`
	Message string `json:"message"`
}
