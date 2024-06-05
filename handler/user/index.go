package user

import (
	"github.com/gin-gonic/gin"
)

type Authorization struct{}

func NewAuthorization(r *gin.RouterGroup) *Authorization {
	h := &Authorization{}

	newRoute(h, r)

	return h
}

func newRoute(h *Authorization, r *gin.RouterGroup) {
	r.GET("/line", h.LineLogin)
	r.GET("/line/callback", h.LineCallback)
}
