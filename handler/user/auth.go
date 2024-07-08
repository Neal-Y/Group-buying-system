package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/config"
	"shopping-cart/constant"
)

func (h *User) LineLogin(c *gin.Context) {
	state := "randomStateString"
	lineURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=profile%%20openid%%20email", constant.LineAuthURL, config.AppConfig.LineClientID, config.AppConfig.LineRedirectURI, state)
	c.Redirect(http.StatusFound, lineURL)
}

func (h *User) LineCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != "randomStateString" {
		handleLineServerError(c, "invalid state")
		return
	}

	user, err := h.service.ExchangeTokenAndGetProfile(code)
	if err != nil {
		handleLineServerError(c, err.Error())
		return
	}

	err = h.service.SaveOrUpdateUser(user)
	if err != nil {
		handleLineServerError(c, "failed to save or update user")
		return
	}

	c.Redirect(http.StatusFound, config.AppConfig.NgrokURL)
}

func handleLineServerError(c *gin.Context, errorMessage string) {
	fmt.Printf("handleLineServerError called with message: %s\n", errorMessage)
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/error?message=%s", config.AppConfig.NgrokURL, errorMessage))
}
