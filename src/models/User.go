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

func CreateUser(client *firestore.Client, user User) error {
	
	exists, err := CheckUserExists(client, user.UserID)
	if err != nil {
		log.Printf("Failed to check if user exists: %v", err)
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	_, _, err = client.Collection("User").Add(context.Background(), user)
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return err
	}

	return nil
}

func CheckUserExists(client *firestore.Client, userID string) (bool, error) {
	query := client.Collection("User").Where("UserID", "==", userID).Documents(context.Background())
	defer query.Stop()

	log.Printf("Checking if user with ID %s exists in Firestore...", userID)


	doc, err := query.Next()
	if err != nil {
		if err == iterator.Done {
		
			log.Printf("No user found with ID %s", userID)
			return false, nil
		}
		log.Printf("Failed to query Firestore: %v", err)
		return false, err
	}


	log.Printf("User found with ID: %v", doc.Ref.ID)
	return true, nil
}

type UserWithID struct {
	IDUser string `json:"DocumentId"`
	Name   string `json:"name"`
	UserID string `json:"userID"`
}

func GetIDUser(client *firestore.Client, UserID string) (UserWithID, error) {
	var userwithid UserWithID

	iter := client.Collection("User").Where("UserID", "==", UserID).Documents(context.Background())
	defer iter.Stop()
	
	doc, err := iter.Next()
	if err != nil {
		log.Printf("Error retrieving document: %v", err)
		return userwithid, err
	}

	if err := doc.DataTo(&userwithid); err != nil {
		log.Printf("Error parsing document data: %v", err)
		return userwithid, err
	}

	userwithid.IDUser = doc.Ref.ID

	return userwithid, nil
}