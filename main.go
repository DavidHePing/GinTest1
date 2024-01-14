package main

import (
	"GinTest1/api/route"
	"GinTest1/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	route.SetUp(gin)
	gin.Run("0.0.0.0:8080")
}
