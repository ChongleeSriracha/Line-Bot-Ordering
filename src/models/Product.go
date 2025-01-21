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


func GetProduct(client *firestore.Client) ([]Product, error) {
	var products []Product

	iter := client.Collection("Product").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err != nil {
			break 
		}

		var product Product
		if err := doc.DataTo(&product); err != nil {
			log.Printf("Error parsing document: %v", err)
			continue 
		}
		products = append(products, product) 
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

		products = append(products, product) 

	}

	return products, nil
}

type ProductID struct {
	ProductID string `json:"DocumentId"`
	Count       int    `json:"count"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Price       int `json:"price"`
	Status      bool   `json:"status"`
	URL         string `json:"url"`
}

func GetProductByName(client *firestore.Client, name string) (ProductID, error) {
	var productid ProductID

	iter := client.Collection("Product").Where("name", "==", name).Documents(context.Background())
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		log.Printf("Error retrieving document: %v", err)
		return productid, err
	}

	if err := doc.DataTo(&productid); err != nil {
		log.Printf("Error parsing document data: %v", err)
		return productid, err
	}
	productid.ProductID = doc.Ref.ID

	return productid, nil
}