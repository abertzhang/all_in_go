package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	router.GET("/asciiJson", func(c *gin.Context) {
		data:=map[string]interface{}{
			"lang": "golang",
			"version": "1.20",
			"msg": "hello, world",
			"tag":"br",
		}

		c.AsciiJSON(http.StatusOK, data) // Using AsciiJSON to return ASCII-encoded JSON
	})
	router.Run(":8080")
}