package route

import (
	"GinTest1/api/middleware"
	"GinTest1/test/controller"
	"GinTest1/test/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetUp(gin *gin.Engine) {
	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.OutputPaths = []string{"./logs/test.log"}
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, _ := config.Build()

	router := gin.Group("")
	SwaggerRouter(router)
	HealthRouter(router)

	apiRouter := gin.Group("api/v1")
	apiRouter.Use(middleware.IpMiddleware())
	controller.NewTestController(usecase.NewTestUseCase(), logger, apiRouter)
}
