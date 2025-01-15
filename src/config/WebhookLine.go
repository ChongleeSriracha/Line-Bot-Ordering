package config

import (
	"line-Bot-Ordering/src/handler"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func WebhookLine() (*linebot.Client, error) {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	chanelSecret := os.Getenv("SECRET_TOKEN")
	chanelAccess := os.Getenv("ACCESS_TOKEN")

	// Initialize LineBot
	bot, err := linebot.New(chanelSecret, chanelAccess)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Webhook client 
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request to get events
		events, err := bot.ParseRequest(r)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		// Call the handler with the events and channelAccessToken
		handler.WebhookHandler(events, bot, chanelAccess)
	})

	return bot, nil
}
