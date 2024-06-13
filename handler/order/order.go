package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
	"strconv"
	"time"
)

func (h *Order) CreateOrder(c *gin.Context) {
	var orderRequest struct {
		UserID       int                    `json:"user_id" binding:"required"`
		Note         string                 `json:"note"`
		OrderDetails []database.OrderDetail `json:"order_details" binding:"required"`
	}

	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := 0.0

	// 檢查庫存跟有效期以及訂單數量不為零
	for i, detail := range orderRequest.OrderDetails {
		if detail.Quantity <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than zero"})
			return
		}

		var product database.Product
		err = infrastructure.Db.First(&product, detail.ProductID).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// 检查库存
		if product.Stock < detail.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product " + product.Name})
			return
		}

		// 检查产品有效期
		if time.Now().After(product.ExpirationTime) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product " + product.Name + " is expired"})
			return
		}

		// order detail 的 price 为 product 的 price
		orderRequest.OrderDetails[i].Price = product.Price

		// 计算总价格
		totalPrice += float64(detail.Quantity) * product.Price
	}

	// 创建订单
	order := database.Order{
		UserID:     orderRequest.UserID,
		TotalPrice: totalPrice,
		Note:       orderRequest.Note,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// 开始事务
	tx := infrastructure.Db.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	for _, detail := range orderRequest.OrderDetails {
		detail.OrderID = order.ID

		// 更新库存
		err = tx.Model(&database.Product{}).Where("id = ?", detail.ProductID).Update("stock", gorm.Expr("stock - ?", detail.Quantity)).Error
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
			return
		}

		err = tx.Create(&detail).Error
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order detail"})
			return
		}
	}

	tx.Commit()

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

	totalPrice := 0.0

	for _, detail := range updateData.OrderDetails {
		var product database.Product
		err := infrastructure.Db.First(&product, detail.ProductID).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// 检查库存
		if product.Stock < detail.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product " + product.Name})
			return
		}

		// 检查产品有效期
		if time.Now().After(product.ExpirationTime) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product " + product.Name + " is expired"})
			return
		}

		// 计算总价格
		totalPrice += float64(detail.Quantity) * product.Price
	}

	updateData.TotalPrice = totalPrice
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
	err := infrastructure.Db.Preload("OrderDetails").First(&order, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// 开始事务
	tx := infrastructure.Db.Begin()

	// 恢复库存
	for _, detail := range order.OrderDetails {
		if err := tx.Model(&database.Product{}).Where("id = ?", detail.ProductID).Update("stock", gorm.Expr("stock + ?", detail.Quantity)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore product stock"})
			return
		}
	}

	// 删除订单
	err = infrastructure.Db.Delete(&order).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	tx.Commit()

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
