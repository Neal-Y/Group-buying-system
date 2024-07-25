package service

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"shopping-cart/config"
)

type NotificationService interface {
	Notify(productID int, productName string, currentStock int)
}

type notificationService struct {
	bot    *linebot.Client
	userID string
}

func NewNotificationService() NotificationService {
	bot, err := linebot.New(config.AppConfig.LineMsgSecret, config.AppConfig.LineMsgToken)
	if err != nil {
		log.Fatal(err)
	}

	userID := config.AppConfig.LineAdminID
	if userID == "" {
		log.Fatal("LINE_USER_ID is not set in the environment variables")
	}

	return &notificationService{
		bot:    bot,
		userID: userID,
	}
}

func (n *notificationService) Notify(productID int, productName string, currentStock int) {
	message := fmt.Sprintf("提醒: 商品 %s (ID: %d) 庫存已低於50%% 目前剩下: %d個單位", productName, productID, currentStock)

	_, err := n.bot.PushMessage(n.userID, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
