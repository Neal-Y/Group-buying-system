package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIDFromPath(c *gin.Context, paramName string) (int, bool) {
	idStr := c.Param(paramName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid " + paramName})
		return 0, false
	}
	return id, true
}
