package main

import (
	"line-Bot-Ordering/src/component"
	"line-Bot-Ordering/src/config"
	"line-Bot-Ordering/src/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize webhook and Firebase connection
	_, err := config.WebhookLine()
	if err != nil {
		log.Fatalf("Failed to initialize webhook Line: %v", err)
	}


	client_firebase , err := config.FirebaseSdk()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase SDK: %v", err)
	}

	// Get ACCESS_TOKEN from environment variable
	channelAccess := os.Getenv("ACCESS_TOKEN")
	if channelAccess == "" {
		log.Fatal("ACCESS_TOKEN environment variable is missing")
	}

	// Create the rich menu for LINE bot
	component.CreateRichMenu(channelAccess)

	// Initialize Gin engine
	r := gin.Default()

	// Register routes for Gin
	routes.RegisterRoutes(r, client_firebase)

	// Start the Gin server in a separate goroutine
	go func() {
		log.Println("Gin server starting at :8081")
		if err := r.Run(":8081"); err != nil {
			log.Fatal("Error starting Gin server: ", err)
		}
	}()

	// Start the LINE webhook server in another goroutine
	go func() {
		log.Println("LINE webhook server starting at :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Error starting LINE server: ", err)
		}
	}()

	// Block the main goroutine to keep servers running
	select {}
}
