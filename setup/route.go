package setup

import (
	"GinTest1/api/controller"
	"GinTest1/api/http"
	"GinTest1/api/middleware"
	"GinTest1/db"
	"GinTest1/repository"
	"GinTest1/usecase"
	"GinTest1/zapLogger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func Setup(gin *gin.Engine) {

	fileLogger := zapLogger.InitLogger()

	router := gin.Group("")
	http.SwaggerRouter(router)
	http.HealthRouter(router)

	apiRouter := gin.Group("api/v1")
	apiRouter.Use(middleware.IpMiddleware())

	Register(apiRouter)

	// controller.NewTestController(usecase.NewTestUseCase(), fileLogger, apiRouter)

	config := ViperSetup()

	testDb := db.Postgre{}.GetGormDb(config.TestDb.GetGormPostgreConnectString(), fileLogger)
	controller.NewCarController(usecase.NewCarUsecase(repository.NewCarRepository(testDb)),
		fileLogger, apiRouter, cache.New(5*time.Minute, 10*time.Minute))

	//GracefulShutdown, must setup after register controller
	GracefulShutdown(gin, config)
}
