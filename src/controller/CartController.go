package controller

import (
	"line-Bot-Ordering/src/models"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func GetCurrentCart(c *gin.Context, client *firestore.Client) {
	userID := c.DefaultQuery("userID", "")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
		return
	}

	userWithID, err := models.GetIDUser(client, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	carts, err := models.GetCurrentCart(client, userWithID)
	if err != nil {
		if err.Error() == "no carts found with current == true" {
			c.JSON(http.StatusNotFound, gin.H{"error": "No current cart found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": carts})
}



func CreateCart(c *gin.Context, client *firestore.Client){

	var cart models.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := models .CreateCart(client, cart)
	if err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart created successfully"})

}

func UpdateProductInCart(c *gin.Context, client *firestore.Client) {
	var request struct {
		Name   string `json:"Name"`
		UserID string `json:"UserID"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
		return
	}

	userWithID, err := models.GetIDUser(client, request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	productID, err := models.GetProductByName(client, request.Name)
	if err != nil {
		log.Printf("Failed to get product by name: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		return
	}

	err = models.UpdateProductInCart(client, userWithID, productID)
	if err != nil {
		log.Printf("Failed to update product in cart: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}