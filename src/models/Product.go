package models

import (
	"context"
	
	"log"

	"cloud.google.com/go/firestore"
)

type Product struct {
	Count       int    `json:"count"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Price       int `json:"price"`
	Status      bool   `json:"status"`
	URL         string `json:"url"`
}

// GetProduct retrieves all products from Firestore
func GetProduct(client *firestore.Client) ([]Product, error) {
	var products []Product

	

	// Use the Firebase client to fetch documents from the 'products' collection
	iter := client.Collection("Product").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err != nil {
			break // Exit the loop when no more documents are available
		}

		var product Product
		if err := doc.DataTo(&product); err != nil {
			log.Printf("Error parsing document: %v", err)
			continue // Skip this document in case of parsing errors
		}
		products = append(products, product) // Add the product to the list
	}

	return products, nil
}


func GetAvaliableProduct(client *firestore.Client)([]Product ,error)  {

	var products []Product 


	iter :=  client.Collection("Product").Where("status", "==", true).Documents(context.Background())

	for {

		doc , err := iter.Next()
		if err != nil {
			break  
		}

		var product Product 
		if err := doc.DataTo(&product); err != nil {
			log.Printf("Error parsing document: %v", err)
			continue 
		}

		products = append(products, product) // Add the product to the list

	}

	return products, nil
}