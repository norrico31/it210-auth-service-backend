package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/google/uuid"
	"github.com/norrico31/it210-auth-service-backend/config"
)

func GenerateToken(userId int) (string, error) {
	exp := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiredAt": time.Now().Add(exp).Unix(),
	})
	secret := []byte(config.Envs.JWTSecret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
