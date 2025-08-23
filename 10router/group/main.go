package main

import (
	"time"

	"github.com/gin-gonic/gin"
)
func main() {
	router:=gin.Default()
	{
		v1 := router.Group("/v1")
		v1.GET("/login", login)
	}
	{
		v2 := router.Group("/v2")
		v2.GET("/login", login)
	}
	router.Run(":8080")
}

func login(ctx *gin.Context) {
	time.Sleep(5 * time.Second) // Simulate a delay for the login process
	// Handle login logic here
	ctx.JSON(200, gin.H{"message": "Login successful"})

}