package setup

import (
	"GinTest1/api/controller"
	"GinTest1/log"
	"GinTest1/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var testControllerSet = wire.NewSet(
	usecase.NewTestUseCase,
	log.InitLogger,
	controller.NewTestController,
)

func InitializeDependencyInjection(apiRouter *gin.RouterGroup) *controller.TestController {
	wire.Build(testControllerSet)
	return &controller.TestController{}
}
