package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"shopping-cart/util"
	"strings"
)

func redirectToLogin(c *gin.Context, expectType string) {
	if expectType == "admin" {
		c.Redirect(http.StatusFound, "/admin/login")
	} else {
		c.Redirect(http.StatusFound, "/api/line/login")
	}
	c.Abort()
}

func JWTAuthMiddleware(expectType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			redirectToLogin(c, expectType)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := util.ParseJWT(tokenString, expectType)
		if err != nil {
			redirectToLogin(c, expectType)
			return
		}

		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			redirectToLogin(c, expectType)
			return
		}

		if userType, ok := mapClaims["type"].(string); ok {
			c.Set("user_type", userType)
		} else {
			redirectToLogin(c, expectType)
			return
		}

		c.Next()
	}
}
