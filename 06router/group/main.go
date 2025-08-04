package main

import (
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
	// Handle login logic here
	ctx.JSON(200, gin.H{"message": "Login successful"})

}