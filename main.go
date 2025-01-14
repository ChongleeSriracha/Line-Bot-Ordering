package main

import (
	"line-Bot-Ordering/src/config"
	"line-Bot-Ordering/src/controller"
	"log"
	"net/http"
	"os"
)

func main() {

	config.WebhookLine()
	config.FirebaseSdk()

	
	channelAccess := os.Getenv("ACCESS_TOKEN")
	if channelAccess == "" {
		log.Fatal("ACCESS_TOKEN environment variable is missing")
	}

	
	controller.CreateRichMenu(channelAccess)

	log.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
