package controller

import (
	"GinTest1/domain"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	testUsecase domain.TestUseCase
}

func NewTestController(testUseCase domain.TestUseCase, router *gin.RouterGroup) {
	controller := &TestController{testUsecase: testUseCase}

	router.GET("/test", controller.getTest)
	router.GET("/test/:id", controller.getTestById)
	router.POST("/test", controller.postTest)
	router.PATCH("/test/:id", controller.patchTest)
	router.DELETE("/test/:id", controller.deleteTest)
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} Test Get
// @Router /test [get]
func (controller *TestController) getTest(ctx *gin.Context) {
	ctx.JSON(200, controller.testUsecase.GetTest())
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Get
// @Router /test/{id} [get]
func (controller *TestController) getTestById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, controller.testUsecase.GetTestById(id))
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} Test Get
// @Router /test [post]
func (controller *TestController) postTest(ctx *gin.Context) {
	ctx.JSON(200, controller.testUsecase.PostTest())
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Get
// @Router /test/{id} [patch]
func (controller *TestController) patchTest(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, controller.testUsecase.PatchTest(id))
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Get
// @Router /test/{id} [delete]
func (controller *TestController) deleteTest(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, controller.testUsecase.DeleteTest(id))
}
