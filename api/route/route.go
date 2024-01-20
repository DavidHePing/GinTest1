package route

import (
	"GinTest1/api/customer_logger"
	"GinTest1/api/middleware"
	"GinTest1/test/controller"
	"GinTest1/test/usecase"

	"github.com/gin-gonic/gin"
)

func SetUp(gin *gin.Engine) {
	file_logger := customer_logger.InitLogger()

	router := gin.Group("")
	SwaggerRouter(router)
	HealthRouter(router)

	apiRouter := gin.Group("api/v1")
	apiRouter.Use(middleware.IpMiddleware())
	controller.NewTestController(usecase.NewTestUseCase(), file_logger, apiRouter)
}
