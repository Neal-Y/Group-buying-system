package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/util"
)

func (h *Product) GetProduct(c *gin.Context) {
	id, _ := util.GetIDFromPath(c, "id")
	product, err := h.productService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *Product) SearchProducts(c *gin.Context) {
	params, err := util.SearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, total, err := h.productService.SearchProducts(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products, "total": total})
}

func (h *Product) GetByID(c *gin.Context) {
	id, _ := util.GetIDFromPath(c, "id")
	product, err := h.productService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
