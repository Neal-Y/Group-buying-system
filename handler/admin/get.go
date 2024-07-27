package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Admin) GetAdmin(c *gin.Context) {
	admin, err := h.adminService.GetAdmin()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"admin": admin})
}
