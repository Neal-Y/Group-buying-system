package general

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func RedirectToLineAuth(c *gin.Context) {
//	// 生成 LINE 授權 URL 並重定向
//	authURL := "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=YOUR_CLIENT_ID&redirect_uri=YOUR_CALLBACK_URL&state=YOUR_STATE&scope=profile"
//	c.Redirect(http.StatusFound, authURL)
//}
//
//func LineAuthCallback(c *gin.Context) {
//	code := c.Query("code")
//	// 使用 code 換取 access_token 並獲取用戶資料
//	// 假設獲取到了用戶資料 user
//	user := getUserFromLine(code)
//	// 建立/更新用戶信息
//	saveOrUpdateUser(user)
//	// 生成 session 或 token
//	token := generateToken(user)
//	c.JSON(http.StatusOK, gin.H{"token": token})
//}
