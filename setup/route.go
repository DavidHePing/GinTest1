package setup

import (
	"GinTest1/api/controller"
	"GinTest1/api/http"
	"GinTest1/api/middleware"
	"GinTest1/db"
	"GinTest1/log"
	"GinTest1/repository"
	"GinTest1/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup(gin *gin.Engine) {
	fileLogger := log.InitLogger()

	router := gin.Group("")
	http.SwaggerRouter(router)
	http.HealthRouter(router)

	apiRouter := gin.Group("api/v1")
	apiRouter.Use(middleware.IpMiddleware())
	controller.NewTestController(usecase.NewTestUseCase(), fileLogger, apiRouter)

	ViperSetup()
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	connection := fmt.Sprint("host=", dbHost, " port=", dbPort, " user=", dbUser, " password=", dbPass, " dbname=", dbName)
	testDb := db.Postgre{}.GetGormDb(connection, fileLogger)
	controller.NewCarController(usecase.NewCarUsecase(*repository.NewCarRepository(testDb)), fileLogger, apiRouter)
}
