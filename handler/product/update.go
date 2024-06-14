package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer"
)

func (h *Product) UpdateProduct(c *gin.Context) {
	id, _ := GetIDFromPath(c, "id")

	var productDto datatransfer.ProductPayload

	err := c.ShouldBindJSON(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.productService.UpdateProduct(id, &productDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product_repository": product})
}
