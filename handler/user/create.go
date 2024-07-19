package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer/user"
)

func (h *User) CreateUser(c *gin.Context) {
	var req user.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}
