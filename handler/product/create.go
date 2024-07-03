package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	product2 "shopping-cart/model/datatransfer/product"
)

func (h *Product) CreateProduct(c *gin.Context) {
	var productDto product2.Payload

	err := c.ShouldBindJSON(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.productService.CreateProduct(&productDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}
