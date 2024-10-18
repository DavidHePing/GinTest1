package setup

import (
	"GinTest1/api/controller"
	"GinTest1/usecase"
	"GinTest1/zapLogger"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func Register(apiRouter *gin.RouterGroup) {
	container := dig.New()
	container.Provide(usecase.NewTestUseCase)
	container.Provide(zapLogger.InitLogger)
	container.Provide(func() *gin.RouterGroup {
		return apiRouter
	})

	err := container.Invoke(func(testUseCase usecase.TestUseCase, logger *zap.Logger, router *gin.RouterGroup) {
		controller.NewTestController(testUseCase, logger, router)
	})

	if err != nil {
		log.Println("Error invoking testController:", err)
	}
}
