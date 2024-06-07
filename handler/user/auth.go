package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"shopping-cart/config"
	"shopping-cart/constant"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

func (h *Authorization) LineLogin(c *gin.Context) {
	state := "randomStateString" // 应该生成随机的state并且保存以验证回调
	lineURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=profile%%20openid%%20email", constant.LineAuthURL, config.AppConfig.LineClientID, config.AppConfig.LineRedirectURI, state)
	c.Redirect(http.StatusFound, lineURL)
}

func (h *Authorization) LineCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != "randomStateString" { // 应该检查保存的state以验证请求
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid state"})
		return
	}

	// 获取access token
	resp, err := http.PostForm(constant.LineTokenURL, url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {config.AppConfig.LineRedirectURI},
		"client_id":     {config.AppConfig.LineClientID},
		"client_secret": {config.AppConfig.LineClientSecret},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get token"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read token response"})
		return
	}

	var tokenData map[string]interface{}
	if err := json.Unmarshal(body, &tokenData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse token response"})
		return
	}

	accessToken := tokenData["access_token"].(string)

	// 获取用户信息
	req, err := http.NewRequest("GET", constant.LineProfileURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create profile request"})
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get profile"})
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read profile response"})
		return
	}

	var profileData struct {
		UserID      string `json:"userId"`
		DisplayName string `json:"displayName"`
		Email       string `json:"email"`
	}

	if err := json.Unmarshal(body, &profileData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse profile response"})
		return
	}

	// 将用户信息保存到数据库
	user := database.User{
		LineID:      profileData.UserID,
		DisplayName: profileData.DisplayName,
		Email:       profileData.Email,
		LineToken:   accessToken,
	}

	err = infrastructure.Db.Where(database.User{LineID: profileData.UserID}).Assign(user).FirstOrCreate(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save or update user"})
		return
	}

	// 设置session或token并返回用户信息
	c.Redirect(http.StatusFound, config.AppConfig.NgrokURL)
}
