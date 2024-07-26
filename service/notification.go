package service

import "fmt"

type NotificationService interface {
	Notify(productID int, productName string, currentStock int)
}

type notificationService struct{}

func NewNotificationService() NotificationService {
	return &notificationService{}
}

func (n *notificationService) Notify(productID int, productName string, currentStock int) {
	fmt.Printf("Alert: Product %s (ID: %d) stock is below 50%% of the initial stock. Current stock: %d\n", productName, productID, currentStock)
}
