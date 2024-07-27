package service

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"shopping-cart/config"
	"shopping-cart/util"
)

type NotificationService interface {
	Notify(userID string, message string) error
	SendEmail(to, subject, body string) error
}

type notificationService struct {
	bot *linebot.Client
}

func NewNotificationService() NotificationService {
	bot, err := linebot.New(config.AppConfig.LineMsgSecret, config.AppConfig.LineMsgToken)
	if err != nil {
		log.Fatal(err)
	}

	return &notificationService{
		bot: bot,
	}
}

func (n *notificationService) Notify(userID string, message string) error {
	_, err := n.bot.PushMessage(userID, linebot.NewTextMessage(message)).Do()
	if err != nil {
		return err
	}
	return nil
}

func (n *notificationService) SendEmail(to, subject, body string) error {
	return util.SendEmail(to, subject, body)
}
