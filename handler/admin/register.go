package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer/admin"
)

func (h *Admin) Register(c *gin.Context) {
	var req admin.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.adminService.RegisterAdmin(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "admin registered successfully"})
}
