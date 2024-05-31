package order

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func CreateOrder(c *gin.Context) {
//	var order models.Order
//	if err := c.BindJSON(&order); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	if err := models.CreateOrder(&order); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, order)
//}
//
//func GetOrderCount(c *gin.Context) {
//	userId := c.Query("user_id")
//	count, err := models.GetOrderCount(userId)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"count": count})
//}
//
//// 其他訂單操作 API
