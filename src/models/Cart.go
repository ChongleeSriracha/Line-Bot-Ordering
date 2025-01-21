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

func GetCurrentCart(client *firestore.Client, user UserWithID) ([]Cart, []string, error) {
	var carts []Cart
	var cartDocIDs []string

	fmt.Printf("iduser = %s\n", user.IDUser)
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
			return nil, nil, err
		}

		var cart Cart
		if err := doc.DataTo(&cart); err != nil {
			log.Printf("Error parsing document data: %v", err)
			continue
		}

		
		carts = append(carts, cart)
		cartDocIDs = append(cartDocIDs, doc.Ref.ID)
		fmt.Println(cart)
		fmt.Println("=================================")
		fmt.Println(cartDocIDs)

	}

	if len(carts) == 0 {
		log.Println("No current cart found")
		return nil, nil, errors.New("no carts found with current == true")
	}

	return carts, cartDocIDs, nil
}

func UpdateProductInCart(client *firestore.Client, user UserWithID, product ProductID) error {

	currentCart, cartDocIDs, err := GetCurrentCart(client, user)
	if err != nil && err.Error() != "no carts found with current == true" {
		log.Printf("Failed to get current cart: %v", err)
		return err
	}

	var cart *Cart
	var cartDocID string

	if len(currentCart) == 0 {
		
		log.Println("No current cart found. Creating a new cart...")
		cart = &Cart{
			Product:  make(map[string]CartDetails),
			Count:    0,
			Price:    0,
			User:     user.IDUser,
			Current:  true,
			UpdateAt: time.Now(),
		}

		productID := product.ProductID
		productPrice := product.Price
		cart.Product[productID] = CartDetails{Price: productPrice, Quantity: 1}
		cart.Count += 1
		cart.Price += productPrice

		
		docRef, _, err := client.Collection("Cart").Add(context.Background(), cart)
		if err != nil {
			log.Printf("Failed to create new cart: %v", err)
			return err
		}
		cartDocID = docRef.ID
		log.Printf("New cart created successfully with ID: %s", cartDocID)
	} else {
		
		cart = &currentCart[0]
		cartDocID = cartDocIDs[0] 

		
		productID := product.ProductID
		productPrice := product.Price
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

	
		_, err = client.Collection("Cart").Doc(cartDocID).Set(context.Background(), cart)
		if err != nil {
			log.Printf("Failed to update cart: %v", err)
			return err
		}
		log.Printf("Cart updated successfully with ID: %s", cartDocID)
	}

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

// validateCart checks and corrects the cart's total price and count.
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
		fmt.Printf("Cart corrected to valid data.\n")
	}
}
