package handler

import (
	"line-Bot-Ordering/src/controller"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func WebhookHandler(events []*linebot.Event, bot *linebot.Client, channelAccessToken string) {
	for _, event := range events {
		userID := event.Source.UserID

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				action := message.Text 
				controller.HandleEventAction(action, userID, channelAccessToken)
			}
		} else if event.Type == linebot.EventTypePostback {
			handlePostbackEvent(event, bot, userID)
		} else {
			log.Printf("Unhandled event type: %v", event.Type)
		}
	}
}

func handlePostbackEvent(event *linebot.Event, bot *linebot.Client, userID string) {
	panic("unimplemented")
}
