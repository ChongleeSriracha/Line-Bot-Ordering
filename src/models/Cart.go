package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

)

type Cart struct {
	Product  map[string]CartDetails `json:"Product"`
	Count    int                    `json:"Count"`
	Price    int                    `json:"Price"`
	User     string                 `json:"User"`
	Current  bool                   `json:"Current"`
	UpdateAt time.Time              `json:"UpdateAt"`
}

type CartDetails struct {
	Price    int `json:"Price"`
	Quantity int `json:"Quantity"`
}

func GetCurrentCart(client *firestore.Client, user UserWithID) ([]Cart, error) {
	var carts []Cart

	fmt.Println("iduser = %s",user.IDUser)
	iter := client.Collection("Cart").
			Where("Current", "==", true).
			Where("User", "==", user.IDUser).
			OrderBy("UpdateAt", firestore.Desc).
			Limit(1).
			Documents(context.Background())
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error retrieving document: %v", err)
			return nil, err
		}

		var cart Cart
		if err := doc.DataTo(&cart); err != nil {
			log.Printf("Error parsing document data: %v", err)
			continue
		}

		carts = append(carts, cart)
	}

	if len(carts) == 0 {
		log.Println("No current cart found")
		return nil, errors.New("no carts found with current == true")
	}

	return carts, nil
}

func UpdateProductInCart(client *firestore.Client, user UserWithID, product ProductID) error {
	currentCart, err := GetCurrentCart(client, user)
	if err != nil {
		log.Printf("Failed to get current cart: %v", err)
		return err
	}

	if len(currentCart) == 0 {
		return errors.New("No current cart found")
	}

	cart := &currentCart[0]
	productID := product.ProductID
	productPrice := product.Price

	
	if cart.Product == nil {
		cart.Product = make(map[string]CartDetails)
	}

	if details, exists := cart.Product[productID]; exists {
		details.Quantity += 1
		cart.Product[productID] = details
		cart.Count += 1
		cart.Price += productPrice
	} else {
		cart.Product[productID] = CartDetails{Price: productPrice, Quantity: 1}
		cart.Count += 1
		cart.Price += productPrice
	}


	validateCart(cart)

	_, err = client.Collection("Cart").Doc(user.UserID).Set(context.Background(), cart)
	if err != nil {
		log.Printf("Failed to update cart: %v", err)
		return err
	}

	log.Printf("Cart updated successfully: %+v", cart)
	return nil
}


func CreateCart(client *firestore.Client, cart Cart) error {

	_, _, err := client.Collection("Cart").Add(context.Background(), cart)

	if err != nil {
		log.Printf("Failed to add cart: %v", err)
		return err
	}

	return nil

}

func validateCart(cart *Cart) {
	var totalPrice, totalCount int

	for _, product := range cart.Product {
		totalPrice += product.Price * product.Quantity
		totalCount += product.Quantity
	}

	if totalPrice == cart.Price && totalCount == cart.Count {
		fmt.Println("Cart validation passed: Final price and count are correct.")
	} else {
		fmt.Printf("Cart validation failed:\n")
		fmt.Printf("Expected Price: %d, Actual Price: %d\n", totalPrice, cart.Price)
		fmt.Printf("Expected Count: %d, Actual Count: %d\n", totalCount, cart.Count)
		cart.Price = totalPrice
		cart.Count = totalCount
		fmt.Printf("change to correct to data:\n")

	}
}
