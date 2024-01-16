package route

import (
	"GinTest1/api/middleware"
	"GinTest1/test/controller"
	"GinTest1/test/usecase"

	"github.com/gin-gonic/gin"
)

func SetUp(gin *gin.Engine) {
	router := gin.Group("")
	SwaggerRouter(router)
	HealthRouter(router)

	apiRouter := gin.Group("api/v1")
	apiRouter.Use(middleware.IpMiddleware())
	controller.NewTestController(usecase.NewTestUseCase(), apiRouter)
}
