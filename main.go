package main

import (
	"line-Bot-Ordering/src/config"

	"log"
	"net/http"
)

func main()  {

	config.WebhookLine()
	config.FirebaseSdk()
	
	log.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}