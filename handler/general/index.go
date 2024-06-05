package general

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type General struct {
}

type Response struct {
	Message string `json:"message"`
}

func NewGeneral(r *gin.RouterGroup) *General {
	h := &General{}

	newRoute(h, r)

	return h
}

func newRoute(h *General, r *gin.RouterGroup) {
	r.GET("/ping", h.ping)
	r.GET("/health", h.health)
	r.GET("/ready", h.ready)
	r.GET("/", h.test)
}

func (h *General) test(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *General) ping(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "pong"})
}

func (h *General) health(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "OK"})
}

// ready
func (h *General) ready(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "OK"})
}
