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

// FirebaseSdk initializes the Firebase SDK and returns a Firestore client
func FirebaseSdk() (*firestore.Client, error) {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	projectID := os.Getenv("PROJECT_ID")
	credentialsFile := os.Getenv("CREDENTIALS_FILE")

	// Initialize Firebase App
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	sa := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
		return nil, err
	}

	// Initialize Firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
		return nil, err
	}

	log.Println("Firebase SDK initialized successfully!")
	return client, nil
}
