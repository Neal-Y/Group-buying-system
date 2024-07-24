package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"shopping-cart/config"
	"shopping-cart/model/datatransfer/user"
)

func (h *User) EmailLogin(c *gin.Context) {
	var req user.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	redirectURL := fmt.Sprintf("%s/api/buffer?token=%s&display_name=%s", config.AppConfig.NgrokURL, token, url.QueryEscape(req.DisplayName))
	c.Redirect(http.StatusFound, redirectURL)
}

func (h *User) EmailRegister(c *gin.Context) {
	var req user.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.RegisterUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}
