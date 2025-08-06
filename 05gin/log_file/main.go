package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)
func main() {
	// gin.DisableConsoleColor()
	f,_:=os.Create("gin.log")
	fmt.Println("Log file created at:", f.Name())
	gin.DefaultWriter = f
	gin.DefaultErrorWriter = f
	router:=gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong123")
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
	// gin.DefaultWriter = os.Stdout // Reset to default writer if needed
	// gin.DefaultErrorWriter = os.Stderr // Reset error writer if needed
}
