package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer/order"
	"shopping-cart/util"
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
	params, err := util.ParseProductSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, total, err := h.orderService.SearchOrders(params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders, "total": total})
}
