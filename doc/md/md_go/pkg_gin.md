## 导包依赖

```go
go get -u github.com/gin-gonic/gin
```



```go
import (
	"github.com/gin-gonic/gin"
)
```

## 启动服务

```go
func main() {
	ginServer := gin.Default()
	ginServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	ginServer.POST("/user/add", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(200, gin.H{
			"message": username + " " + password,
		})
	})
	err := ginServer.Run(":80")
	if err != nil {
		return
	}
}
```

## SSL

```
package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
 
func main() {
    router := gin.New()
 
    router.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "success",
        })
    })
 
    // 可以直接用
    //router.RunTLS("0.0.0.0:10679", "./certs/server.cer", "./certs/server.key")
    server := &http.Server{Addr: "0.0.0.0:10679", Handler: router}
    _ = server.ListenAndServeTLS("./certs/server.cer", "./certs/server.key")
}
```

