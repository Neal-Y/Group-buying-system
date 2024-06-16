package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/config"
	"shopping-cart/constant"
)

func (h *Authorization) LineLogin(c *gin.Context) {
	state := "randomStateString"
	lineURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=profile%%20openid%%20email", constant.LineAuthURL, config.AppConfig.LineClientID, config.AppConfig.LineRedirectURI, state)
	c.Redirect(http.StatusFound, lineURL)
}

func (h *Authorization) LineCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != "randomStateString" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid state"})
		return
	}

	accessToken, err := h.service.ExchangeToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange token"})
		return
	}

	user, err := h.service.GetLineProfile(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get profile"})
		return
	}

	err = h.service.SaveOrUpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save or update user"})
		return
	}

	c.Redirect(http.StatusFound, config.AppConfig.NgrokURL)
}
