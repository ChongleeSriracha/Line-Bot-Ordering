package config

import(
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func WebhookLine()  {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}


	// Get environment variables
	chanelSecret := os.Getenv("SECRET_TOKEN")
	chanelAcess := os.Getenv("ACCESS_TOKEN")

	// Initialize LineBot
	bot,err := linebot.New(chanelSecret,chanelAcess)

	if err != nil {
		log.Fatal(err)
	}

	// Initialize Webhook client 
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		_ ,err := bot.ParseRequest(r)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		// events, err := bot.ParseRequest(r)
		// for _, event := range events {
		// 	if event.Type == linebot.EventTypeMessage {
		// 		switch message := event.Message.(type) {
		// 		case *linebot.TextMessage:
		// 			if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("คุณส่งข้อความ: "+message.Text)).Do(); err != nil {
		// 				log.Print(err)
		// 			}
		// 		}
		// 	}
		// }
	})
	
}