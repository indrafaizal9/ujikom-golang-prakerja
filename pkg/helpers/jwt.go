package helpers

import (
	"time"
	"ujikom/config"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id uint, email string) (string, error) {
	const TokenDuration = time.Hour * 24

	appConfig, err := config.LoadConfig(".")
	if err != nil {
		return "", err
	}

	claim := jwt.MapClaims{
		"id":       id,
		"username": email,
		"iss":      "ujikom",
		"iat":      jwt.TimeFunc().Unix(),
		"exp":      jwt.TimeFunc().Add(TokenDuration).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := parseToken.SignedString([]byte(appConfig.AppSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
