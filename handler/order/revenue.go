package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/util"
)

func (h *Order) GetRevenue(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	timezone := c.DefaultQuery("timezone", "UTC")

	startDateUTC, endDateUTC, err := util.ConvertDateRangeToUTC(startDateStr, endDateStr, timezone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	revenue, err := h.orderService.GetRevenueByTimePeriod(startDateUTC, endDateUTC)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"revenue": revenue})
}
