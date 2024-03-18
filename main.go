package main

import (
	"GinTest1/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	setup.Setup(gin)

	setup.GracefulShutdown(gin)
}
