package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"shopping-cart/config"
	"time"
)

var jwtSecret = []byte(config.AppConfig.Secret)

func exportJWTMapClaims(typeToken string) jwt.MapClaims {
	return jwt.MapClaims{
		"type": typeToken,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
}

func GenerateJWT(typeToken string) (string, error) {
	claims := exportJWTMapClaims(typeToken)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string, expectedType string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != expectedType {
		return nil, fmt.Errorf("invalid token type")
	}

	return claims, nil
}
