package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
	"strconv"
)

// fixme:檢查product stock是否滿足order detail 的quantity，夠了才可以下訂單

func (h *Order) CreateOrder(c *gin.Context) {
	var order database.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := order.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 重新加载订单以包含关联的数据
	err := order.FindByID(order.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reload order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
}

func (h *Order) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order := &database.Order{}
	err = order.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (h *Order) UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var updateData database.Order
	err = c.ShouldBindJSON(&updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := &database.Order{}
	err = order.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	updateData.ID = id // 确保更新的ID正确
	err = order.Update(&updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 重新加载订单信息
	err = order.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reload order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": order})
}

func (h *Order) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order database.Order

	// 查找订单
	if err := infrastructure.Db.Preload("OrderDetails").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// 删除关联的订单详情
	if err := infrastructure.Db.Delete(&order.OrderDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order details"})
		return
	}

	// 删除订单
	if err := infrastructure.Db.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func (h *Order) ListOrders(c *gin.Context) {
	orders, err := database.FindAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
