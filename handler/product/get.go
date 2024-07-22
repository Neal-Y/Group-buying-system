package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/util"
	"strconv"
	"time"
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
	keyword := c.Query("keyword")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date"})
			return
		}
	}
	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date"})
			return
		}
	}

	products, total, err := h.productService.SearchProducts(keyword, startDate, endDate, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products, "total": total})
}
