package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pill_dis/envloader"
)

func init() {
	envloader.Loadenv()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
