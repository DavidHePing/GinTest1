package controller

import (
	"GinTest1/domain"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func NewTestController(testUseCase domain.TestUseCase, router *gin.RouterGroup) {
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, testUseCase.GetTest())
	})

	router.GET("/test/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, testUseCase.GetTestById(id))
	})

	router.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, testUseCase.PostTest())
	})

	router.PATCH("/test/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, testUseCase.PatchTest(id))
	})

	router.DELETE("/test/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, testUseCase.DeleteTest(id))
	})
}
