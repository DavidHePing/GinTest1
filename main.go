package main

import (
	"GinTest1/api/route"
	"GinTest1/graceful"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	route.SetUp(gin)

	graceful.GracefulShutdown(gin)
}
