package controller

import (
	"net/http"
	"line-Bot-Ordering/src/models"

	"github.com/gin-gonic/gin"
)

// GetAllProducts handles the request to retrieve all products
func GetAllProducts(c *gin.Context) {
	products, err := models.GetProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}


func GetAvaliableProducts(c *gin.Context) {
	products, err := models.GetAvaliableProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
