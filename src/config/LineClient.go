package config

import (
    "line-Bot-Ordering/src/services"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/line/line-bot-sdk-go/v7/linebot"
)

func WebhookLine() (*linebot.Client, string, error) {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    channelSecret := os.Getenv("SECRET_TOKEN")
    channelAccess := os.Getenv("ACCESS_TOKEN")

    bot, err := linebot.New(channelSecret, channelAccess)
    if err != nil {
        log.Fatal(err)
    }
    services.CreateRichMenu(channelAccess)

    return bot, channelAccess, nil
}