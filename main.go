package main

import (
	"GinTest1/api/http"
	"GinTest1/api/middleware"
	"GinTest1/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()

	router := gin.Group("")
	http.SwaggerRouter(router)
	http.HealthRouter(router)

	apiRouter := gin.Group("api/v1")
	apiRouter.Use(middleware.IpMiddleware())

	config := setup.ViperSetup()
	setup.Register(apiRouter, &config)

	//GracefulShutdown, must setup after register controller
	setup.GracefulShutdown(gin, config)
}
