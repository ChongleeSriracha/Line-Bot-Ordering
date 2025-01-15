package routes

import (
    "cloud.google.com/go/firestore"
    "github.com/gin-gonic/gin"
    "line-Bot-Ordering/src/controller"
)

func RegisterRoutes(r *gin.Engine, client *firestore.Client) {

    api := r.Group("/api")
    {
        // test api
        api.GET("/test", func(c *gin.Context) {
            c.JSON(200, gin.H{"message": "Server is up"})
        })

        // Product
        api.GET("/products", func(c *gin.Context) {
            controller.GetAllProducts(c, client)
        })
        api.GET("/products/avaliable", func(c *gin.Context) {
            controller.GetAvaliableProducts(c, client)
        })

		
	
    }


}