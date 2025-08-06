package main

import "github.com/gin-gonic/gin"
func main() {
	r:= gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	auth:=r.Group("/")
	auth.Use(AuthMiddleware())
	{
		auth.GET("/hello", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"msg":"hello",
			})
		})
		// 嵌套路由组
		test:= auth.Group("test")
		{
			test.GET("/world", func(c *gin.Context) {
				c.JSON(200,gin.H{
					"msg":"world-test",
				})
			})
		}
	}
	// 监听
	r.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token:=c.GetHeader("token")
		if token!=""{
			c.Next()
		}else {
			c.JSON(401,gin.H{
				"msg":"unauthorized",
			})
			c.Next()
			// c.Abort()
		}
	}}
	
