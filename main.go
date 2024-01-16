package main

import (
	"GinTest1/api/graceful"
	"GinTest1/api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	route.SetUp(gin)

	graceful.GracefulShutdown(gin)
}
