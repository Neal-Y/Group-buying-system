package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/util"
)

func (h *Order) GetRevenue(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startTime, err := util.ValidateTime(startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endTime, err := util.ValidateTime(endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	revenue, err := h.orderService.GetRevenueByTimePeriod(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"revenue": revenue})
}
