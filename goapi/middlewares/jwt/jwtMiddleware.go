package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"louissantucci/goapi/responses"
	"net/http"
)

func JWTTokenCheck(c *gin.Context) {
	jwtToken, err := ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadGateway, err.Error()))
		return
	}

	token, err := ParseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		errorData := "Cannot parse MapClaims"
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, errorData))
		return
	}

	c.Next()
}
