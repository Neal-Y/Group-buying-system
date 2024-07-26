package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LineClientID     string
	LineClientSecret string
	NgrokURL         string
	LineRedirectURI  string
	Secret           string
	Gmail            string
	GmailSecret      string
	LineMsgID        string
	LineMsgSecret    string
	LineMsgToken     string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	AppConfig = Config{
		LineClientID:     viper.GetString("LINE_CLIENT_ID"),
		LineClientSecret: viper.GetString("LINE_CLIENT_SECRET"),
		NgrokURL:         viper.GetString("NGROK_URL"),
		LineRedirectURI:  viper.GetString("NGROK_URL") + "/api/line/callback",
		Secret:           viper.GetString("SECRET"),
		Gmail:            viper.GetString("GMAIL"),
		GmailSecret:      viper.GetString("GMAIL_SECRET"),
		LineMsgID:        viper.GetString("LINE_MSG_ID"),
		LineMsgSecret:    viper.GetString("LINE_MSG_SECRET"),
		LineMsgToken:     viper.GetString("LINE_MSG_TOKEN"),
	}
}
