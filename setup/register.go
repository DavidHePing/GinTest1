package setup

import (
	"GinTest1/api/controller"
	"GinTest1/db"
	"GinTest1/domain"
	"GinTest1/repository"
	"GinTest1/usecase"
	"GinTest1/zapLogger"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Register(apiRouter *gin.RouterGroup, config *domain.Config) {
	fileLogger := zapLogger.InitLogger()
	container := dig.New()
	container.Provide(usecase.NewTestUseCase)
	container.Provide(func() *zap.Logger {
		return fileLogger
	})
	container.Provide(func() *gin.RouterGroup {
		return apiRouter
	})
	container.Provide(func() *gorm.DB {
		return db.Postgre{}.GetGormDb(config.TestDb.GetGormPostgreConnectString(), fileLogger)
	})
	container.Provide(func() *cache.Cache {
		return cache.New(5*time.Minute, 10*time.Minute)
	})
	container.Provide(repository.NewCarRepository)
	container.Provide(usecase.NewCarUsecase)

	err := container.Invoke(func(testUseCase usecase.TestUseCase, logger *zap.Logger, router *gin.RouterGroup) {
		controller.NewTestController(testUseCase, logger, router)
	})

	if err != nil {
		log.Println("Error invoking testController:", err)
	}

	err = container.Invoke(func(carUseCase usecase.CarUseCase, logger *zap.Logger, router *gin.RouterGroup, cache *cache.Cache) {
		controller.NewCarController(carUseCase, logger, router, cache)
	})

	if err != nil {
		log.Println("Error invoking carController:", err)
	}
}
