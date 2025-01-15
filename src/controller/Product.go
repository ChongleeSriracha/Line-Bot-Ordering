package controller

import (
	"line-Bot-Ordering/src/models"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

// GetAllProducts handles the request to retrieve all products
func GetAllProducts(c *gin.Context, client *firestore.Client) {
	products, err := models.GetProduct(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}


func GetAvaliableProducts(c *gin.Context, client *firestore.Client) {
	products, err := models.GetAvaliableProduct(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
