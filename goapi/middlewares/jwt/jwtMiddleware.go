package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"louissantucci/goapi/config"
	"louissantucci/goapi/models"
	"net/http"
	"strings"
)

func GetSecretKey() string {
	return config.Jwt_secret
}

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("no header given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrect Authorization header format")
	}

	return jwtToken[1], nil
}

func ParseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, OK := token.Method.(*jwt.SigningMethodHMAC)
		if !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(GetSecretKey()), nil
	})

	if err != nil {
		return nil, errors.New("bad JWT token")
	}

	return token, nil
}

func JWTTokenCheck(c *gin.Context) {
	jwtToken, err := ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := ParseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	c.Next()
}
