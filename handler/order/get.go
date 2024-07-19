package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/datatransfer/order"
	"shopping-cart/util"
)

func (h *Order) GetOrder(c *gin.Context) {
	id, err := util.GetIDFromPath(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (h *Order) ListOrders(c *gin.Context) {
	orders, err := h.orderService.ListAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

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

func (h *Order) GetOrderByUserDisplayName(c *gin.Context) {
	var orderRequest order.GetOrderByUsername
	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.orderService.GetOrderByUserDisplayName(orderRequest.DisplayName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "display_name not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
