package main

import (
	"GinTest1/route"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	route.SetUp(gin)
	gin.Run("0.0.0.0:8080")
}
