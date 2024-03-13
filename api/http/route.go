package http

import (
	"GinTest1/api/controller"
	"GinTest1/api/middleware"
	"GinTest1/log"
	"GinTest1/usecase"

	"github.com/gin-gonic/gin"
)

func SetUp(gin *gin.Engine) {
	fileLogger := log.InitLogger()

	router := gin.Group("")
	SwaggerRouter(router)
	HealthRouter(router)

	apiRouter := gin.Group("api/v1")
	apiRouter.Use(middleware.IpMiddleware())
	controller.NewTestController(usecase.NewTestUseCase(), fileLogger, apiRouter)
}
