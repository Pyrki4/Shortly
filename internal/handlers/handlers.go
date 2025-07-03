package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateShortLink(c *gin.Context) {
	url := c.Request.Form.Get("url")

	fmt.Println(url)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
