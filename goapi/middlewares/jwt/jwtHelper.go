package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"louissantucci/goapi/config"
	"louissantucci/goapi/models"
	"strings"
	"time"
)

var jwtKey = []byte(config.Jwt_secret)

const EXPIRATION_TIME = 1 * time.Hour

type Claim struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	jwt.RegisteredClaims
}

func GetSecretKey() string {
	return config.Jwt_secret
}

func GenerateJWT(user *models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(EXPIRATION_TIME)
	claims := &Claim{
		Id:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func GetClaimFromHeader(header string) (*Claim, error) {
	tokenStr, err := ExtractBearerToken(header)
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	return claims, nil
}

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("You need to be logged in to perform this action.")
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
