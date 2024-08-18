//go:build wireinject

package setup

import (
	"GinTest1/api/controller"
	"GinTest1/domain"
	"GinTest1/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitializeController(logger *zap.Logger, router *gin.RouterGroup) *controller.TestController {
	wire.Build(
		controller.NewTestController,
		usecase.NewTestUseCase,
		wire.Bind(new(domain.TestUseCase), new(*usecase.TestUsecase)),
	)
	return nil
}
