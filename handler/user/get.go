package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/util"
	"strconv"
)

func (h *User) GetUser(c *gin.Context) {
	id, err := util.GetIDFromPath(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *User) SearchUsers(c *gin.Context) {
	params, err := util.SearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isMember, _ := strconv.ParseBool(c.Query("is_member"))

	users, total, err := h.userService.SearchUsers(params, isMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "total": total})
}
