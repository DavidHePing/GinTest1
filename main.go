package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	router := gin.New()

	router.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World")
	})

	router.Run("0.0.0.0:8080")
}
