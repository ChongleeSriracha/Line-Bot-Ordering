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
        api.GET("/product", func(c *gin.Context) {  
            controller.GetIDProduct(c, client)
        })

        // User
		api.POST("/user", func(c *gin.Context) {        
            controller.CreateUser(c, client)
        })
        api.GET("/user/id", func(c *gin.Context) {
            controller.GetIDUser(c,client)
        } )

        //Cart
        api.GET("/cart", func(c *gin.Context) {
            controller.GetCurrentCart(c, client)
        })

        api.POST("/cart/add", func(c *gin.Context) {
            controller.CreateCart(c, client)
        })

        api.PUT("/cart/update", func(c *gin.Context) {
            controller.UpdateProductInCart(c, client)
        })

	
    }


}