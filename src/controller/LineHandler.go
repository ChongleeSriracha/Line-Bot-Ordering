package controller

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func WebhookLineHandler(channelAccessToken string, bot *linebot.Client) {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		events, err := bot.ParseRequest(r)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		HandleEventData(events, bot, channelAccessToken)

	})
}
