package main

import (
    "line-Bot-Ordering/src/config"
    "line-Bot-Ordering/src/controller"
    "line-Bot-Ordering/src/routes"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize webhook and Firebase connection
    clientLine, channelAccess,err := config.WebhookLine()
    if err != nil {
        log.Fatalf("Failed to initialize webhook Line: %v", err)
    }
    clientFirebase, err := config.FirebaseSdk()
    if err != nil {
        log.Fatalf("Failed to initialize Firebase SDK: %v", err)
    }

    // Initialize Gin engine
    r := gin.Default()
    routes.RegisterRoutes(r, clientFirebase)

    controller.WebhookLineHandler(channelAccess, clientLine)

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

    select {}
}