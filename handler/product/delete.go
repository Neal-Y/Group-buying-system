package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Product) DeleteProduct(c *gin.Context) {
	id, _ := GetIDFromPath(c, "id")

	err := h.productService.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
