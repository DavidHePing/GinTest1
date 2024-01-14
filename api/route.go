package route

import (
	"GinTest1/test/controller"
	"GinTest1/test/usecase"

	"github.com/gin-gonic/gin"
)

func SetUp(gin *gin.Engine) {
	router := gin.Group("")
	SwaggerRouter(router)

	apiRouter := gin.Group("api/v1")
	controller.NewTestController(usecase.NewTestUseCase(), apiRouter)
}
