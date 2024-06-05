package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

// 沒有人會把這重要的東西丟這你這低能！記得refactor
const (
	lineAuthURL      = "https://access.line.me/oauth2/v2.1/authorize"
	lineTokenURL     = "https://api.line.me/oauth2/v2.1/token"
	lineProfileURL   = "https://api.line.me/v2/profile"
	lineClientID     = "2005525590"
	lineClientSecret = "3c4461177d114bd691d5657e15ea8ed2"
	lineRedirectURI  = "https://8739-2001-b011-3808-9748-ec26-2b45-2150-4b6b.ngrok-free.app/api/line/callback"
)

func (h *Authorization) LineLogin(c *gin.Context) {
	state := "randomStateString" // 應該生成隨機的state並且保存以驗證回調
	lineURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=profile", lineAuthURL, lineClientID, lineRedirectURI, state)
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
	resp, err := http.PostForm(lineTokenURL, url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {lineRedirectURI},
		"client_id":     {lineClientID},
		"client_secret": {lineClientSecret},
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
	req, err := http.NewRequest("GET", lineProfileURL, nil)
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

	//err = infrastructure.Db.FirstOrCreate(&user, database.User{LineID: profileData.UserID}).Error

	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
	//	return
	//}

	//err = infrastructure.Db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},                     // key colume
	//	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}), // column needed to be updated
	//}).Create(&users)

	err = infrastructure.Db.First(&user).Error

	if err != nil {
		infrastructure.Db.Create(&user)
	} else {
		infrastructure.Db.Model(&user).Where("Line_id = ?", profileData.UserID).Update("DisplayName", profileData.DisplayName)
	}

	// 设置session或token并返回用户信息
	c.Redirect(http.StatusFound, fmt.Sprintf("https://8739-2001-b011-3808-9748-ec26-2b45-2150-4b6b.ngrok-free.app"))
}
