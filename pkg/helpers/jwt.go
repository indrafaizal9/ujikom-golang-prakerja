package helpers

import (
	"errors"
	"strings"
	"time"
	"ujikom/config"
	"ujikom/pkg/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(model models.User) (string, error) {
	const TokenDuration = time.Hour * 24

	appConfig, err := config.LoadConfig(".")
	if err != nil {
		return "", err
	}

	claim := jwt.MapClaims{
		"id":       model.ID,
		"username": model.Username,
		"role":     model.Role,
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

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Unauthorized")
	headerToken := c.Request.Header.Get("Authorization")
	bearerToken := strings.HasPrefix(headerToken, "Bearer ")

	if !bearerToken {
		return nil, errResponse
	}

	stringsToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringsToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte("secret"), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
