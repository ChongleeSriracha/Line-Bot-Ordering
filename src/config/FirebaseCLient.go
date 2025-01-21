package config

import (
	"context"
	"log"
	"os"

	"firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
)


func FirebaseSdk() (*firestore.Client, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	projectID := os.Getenv("PROJECT_ID")
	credentialsFile := os.Getenv("CREDENTIALS_FILE")

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	sa := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
		return nil, err
	}

	log.Println("Firebase SDK initialized successfully!")
	return client, nil
}
