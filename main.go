package main

import (
	"GinTest1/api/http"
	"GinTest1/graceful"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	http.SetUp(gin)

	graceful.GracefulShutdown(gin)
}
