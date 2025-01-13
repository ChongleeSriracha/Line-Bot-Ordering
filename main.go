package main

import (
	"line-Bot-Ordering/src/config"
	"line-Bot-Ordering/src/view"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize Webhook and Firebase SDK
	config.WebhookLine()
	config.FirebaseSdk()

	// Get channel access token from environment variables
	channelAccess := os.Getenv("ACCESS_TOKEN")
	if channelAccess == "" {
		log.Fatal("ACCESS_TOKEN environment variable is missing")
	}

	// Call RichMenu function from view package and pass the channel access token
	view.CreateRichMenu(channelAccess)

	// Start server
	log.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
