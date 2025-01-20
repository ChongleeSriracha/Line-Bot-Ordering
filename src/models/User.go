package models

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type User struct {
	Name   string `json:"name"`
	UserID string `json:"userID"`
}

// CreateUser adds a new user to the Firestore collection if the user doesn't already exist.
func CreateUser(client *firestore.Client, user User) error {
	// Check if the user already exists
	exists, err := CheckUserExists(client, user.UserID)
	if err != nil {
		log.Printf("Failed to check if user exists: %v", err)
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	// Add the user to the Firestore collection
	_, _, err = client.Collection("User").Add(context.Background(), user)
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return err
	}

	return nil
}

func CheckUserExists(client *firestore.Client, userID string) (bool, error) {
	// Perform the query to check if a user with the given userID exists
	query := client.Collection("User").Where("UserID", "==", userID).Documents(context.Background())
	defer query.Stop()

	// Log the userID being checked
	log.Printf("Checking if user with ID %s exists in Firestore...", userID)

	// Iterate through the results to check if any document matches
	doc, err := query.Next()
	if err != nil {
		if err == iterator.Done {
			// No documents found
			log.Printf("No user found with ID %s", userID)
			return false, nil
		}
		log.Printf("Failed to query Firestore: %v", err)
		return false, err
	}

	// Document found
	log.Printf("User found with ID: %v", doc.Ref.ID)
	return true, nil
}
