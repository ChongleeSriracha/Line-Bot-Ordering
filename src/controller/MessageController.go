package controller

import (
	"line-Bot-Ordering/src/models"
	"line-Bot-Ordering/src/services"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// WebhookHandler handles incoming LINE webhook events
func HandleEventData(events []*linebot.Event, bot *linebot.Client, channelAccessToken string) {
	for _, event := range events {
		userID := event.Source.UserID
		profile, err := bot.GetProfile(userID).Do()
        if err != nil {
            log.Printf("Error getting user profile: %v", err)
            continue
        }
        name := profile.DisplayName

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				action := message.Text
				models.CallCreateUser(userID, name)
				HandleEventAction(action, userID, name,channelAccessToken)
			}
		} else if event.Type == linebot.EventTypePostback {
			handlePostbackEvent(event, bot, userID)
		} else {
			log.Printf("Unhandled event type: %v", event.Type)
		}
	}
}

// handleEventAction processes user actions and sends appropriate responses
func HandleEventAction(action string, userID, name string,channelAccessToken string) {
	if action == "Product" {
		err := services.FlexProduct(userID, channelAccessToken)
		if err != nil {
			log.Fatal("Error creating JSON flex product")
			return
		}
		log.Printf("Flex message sent successfully via push")
	}
	if action == "Order" {
		err := services.OrderProduct(userID , name)
		if err != nil {		
			log.Fatal("Error creating order product")
			return
		}

	}
}

// handlePostbackEvent handles LINE postback events
func handlePostbackEvent(event *linebot.Event, bot *linebot.Client, userID string) {
	panic("unimplemented")
}
