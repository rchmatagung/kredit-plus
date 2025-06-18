package utils

import (
	"boilerplate/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenereateJWT(conf *config.Config, employeeName string, roleName string) (string, error) {
	atSecretKey := conf.Authorization.JWT.AccessTokenSecretKey

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["employee_name"] = employeeName
	claims["role_name"] = roleName
	claims["exp"] = time.Now().Add(time.Hour * 24 * conf.Authorization.JWT.AccessTokenDuration).Unix()
	tokenString, err := token.SignedString([]byte(atSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
