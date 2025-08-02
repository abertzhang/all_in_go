package main

import "github.com/gin-gonic/gin"

func main() {
	ginService := gin.Default()
	ginService.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello,world"})
	})
	err := ginService.Run(":8080")
	if err != nil {
		return
	}
}
