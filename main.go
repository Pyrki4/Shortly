package main

import (
	"github.com/gin-gonic/gin"
	"shortly/internal/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/ping", handlers.CreateShortLink)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
