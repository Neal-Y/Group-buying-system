package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/util"
)

func (h *Admin) ListAdmins(c *gin.Context) {
	admins, err := h.adminService.GetAllAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"admins": admins})
}

func (h *Admin) GetAdmin(c *gin.Context) {
	id, err := util.GetIDFromPath(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin ID"})
		return
	}

	admin, err := h.adminService.GetAdminByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"admin": admin})
}
