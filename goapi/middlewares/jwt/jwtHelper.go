package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"louissantucci/goapi/config"
	"time"
)

var jwtKey = []byte(config.Jwt_secret)

type Claim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func GetEmailFromToken(jwtToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}

	return claims.Email, nil

}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
