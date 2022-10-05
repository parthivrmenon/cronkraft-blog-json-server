package main

import (
	"net/http"

	"parthivrmenon/conkraft-blog-json-server/controller"
	"parthivrmenon/conkraft-blog-json-server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/blogs", func(c *gin.Context) {
		var svc service.BlogService = service.New()
		var handler controller.BlogController = controller.New(svc)
		c.JSON(http.StatusOK, handler.FindAll())
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
