package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)


// Gin web application in Golang, the AuthRequired function returns an anonymous function to define middleware in a flexible and reusable manner. Let’s break down why returning an anonymous function is a common pattern for middleware:
func AuthReq (allowedToken string) gin.HandlerFunc{
	return func (c *gin.Context)  {
		token := c.GetHeader("Authorization")

		if token != allowedToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
		}

		c.Next();
	}
}
func main(){
	r := gin.Default()


	r.GET("/public", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the public endpoint!"})
	})

	// Authorized routes group
	authorized := r.Group("/", AuthReq("my-secret-token"))
	{
		authorized.GET("/secure", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "Welcome to the secure endpoint!"})
        })
        authorized.GET("/profile", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "User profile information"})
        })
	}

	r.Run(":8080") 
}

