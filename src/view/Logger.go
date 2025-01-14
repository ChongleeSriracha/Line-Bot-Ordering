package view

import "log"

func LogMessage(message string) {
	log.Println(message)
}

func LogError(err error, context string) {
	log.Fatalf("%s: %v", context, err)
}
