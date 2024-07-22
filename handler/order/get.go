package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer/order"
	"strconv"
	"time"
)

func (h *Order) ListHistoryOrdersByUser(c *gin.Context) {
	var orderRequest order.ListHistory
	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.orderService.ListHistoryOrdersByDisplayNameAndProductID(orderRequest.DisplayName, orderRequest.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "display_name or product_id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (h *Order) SearchOrders(c *gin.Context) {
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

	orders, total, err := h.orderService.SearchOrders(keyword, startDate, endDate, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders, "total": total})
}
