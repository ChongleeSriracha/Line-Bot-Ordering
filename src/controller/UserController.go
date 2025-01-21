package controller

import (
	"line-Bot-Ordering/src/models"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func  CreateUser(c* gin.Context,client *firestore.Client,) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {	
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}


	err := models.CreateUser(client, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func GetIDUser(c *gin.Context, Client *firestore.Client) {
	UserID := c.DefaultQuery("UserID", "")

	if UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
		return
	}

	userwithid, err := models.GetIDUser(Client, UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userwithid": userwithid})
}

