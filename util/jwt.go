package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"shopping-cart/config"
	"time"
)

var jwtSecret = []byte(config.AppConfig.Secret)

func GenerateJWT(adminID int) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func GenerateResetToken(adminID int) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyResetToken(tokenString string) (int, error) {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return 0, err
	}

	if adminID, ok := claims.(jwt.MapClaims)["admin_id"].(float64); ok {
		return int(adminID), nil
	}

	return 0, fmt.Errorf("invalid token")
}
