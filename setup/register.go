package setup

import (
	"GinTest1/api/controller"
	"GinTest1/log"
	"GinTest1/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Register(apiRouter *gin.RouterGroup) {

	container := dig.New()
	// container.Provide(ViperSetup())
	container.Provide(usecase.NewTestUseCase)
	container.Provide(log.InitLogger)
	container.Provide(apiRouter)
	container.Provide(controller.NewTestController)
}
