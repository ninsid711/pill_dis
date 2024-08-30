package main

import (
	"github.com/gin-gonic/gin"
	"githug.com/ninsid711/pill_dis/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
