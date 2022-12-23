package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"louissantucci/goapi/config"
	"strings"
	"time"
)

var jwtKey = []byte(config.Jwt_secret)

type Claim struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func GetSecretKey() string {
	return config.Jwt_secret
}

func GenerateJWT(email string, name string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claim{
		Email: email,
		Name:  name,
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
