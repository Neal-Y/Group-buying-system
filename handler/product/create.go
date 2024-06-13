package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer"
)

func (h *Product) CreateProduct(c *gin.Context) {
	var productDto datatransfer.ProductPayload

	err := c.ShouldBindJSON(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productDao := database.Product{
		Name:           productDto.Name,
		Picture:        productDto.Picture,
		Price:          productDto.Price,
		Stock:          productDto.Stock,
		Description:    productDto.Description,
		ExpirationTime: productDto.ExpirationTime,
	}

	err = infrastructure.Db.Where("name = ?", productDto.Name).First(&productDao).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name already exists"})
		return
	}

	err = productDao.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": productDao})
}
