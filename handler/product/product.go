package product

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"shopping-cart/model/database"
)

func (h *Product) CreateProduct(c *gin.Context) {
	var product database.Product

	// 使用 JSON 解析请求体
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 解析时间字符串
	if product.ExpirationTime.IsZero() {
		expirationDateStr := c.PostForm("expiration_time")
		expirationDate, err := time.Parse(time.RFC3339, expirationDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiration date"})
			return
		}
		product.ExpirationTime = expirationDate
	}

	// 创建产品
	if err := product.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}

func (h *Product) GetProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product := database.Product{}
	if err := product.FindByID(productID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Product) GetAllProducts(c *gin.Context) {
	products, err := database.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Product) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var updateData database.Product

	// 使用 JSON 解析请求体
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 解析时间字符串
	if expirationDateStr, exists := c.GetPostForm("expiration_time"); exists {
		expirationDate, err := time.Parse(time.RFC3339, expirationDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiration date"})
			return
		}
		updateData.ExpirationTime = expirationDate
	}

	product := &database.Product{}
	if err := product.FindByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := product.Update(&updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": updateData})
}

func (h *Product) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product := database.Product{ID: productID}
	if err := product.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
