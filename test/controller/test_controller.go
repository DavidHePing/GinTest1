package controller

import (
	"GinTest1/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestController struct {
	testUsecase domain.TestUseCase
	logger      *zap.Logger
}

func NewTestController(testUseCase domain.TestUseCase, logger *zap.Logger, router *gin.RouterGroup) {
	controller := &TestController{testUsecase: testUseCase, logger: logger}

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
	ctx.JSON(http.StatusOK, controller.testUsecase.GetTest())
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Get
// @Router /test/{id} [get]
func (controller *TestController) getTestById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad Request!!!"})
		return
	}

	headerIp := ctx.Request.Header.Get("X-Request-Id")
	headerIp2 := ctx.Request.Header.Get("X-Customer-Header")
	log.Println("header ip:", headerIp, "customer header:", headerIp2)

	defer controller.logger.Sync()
	controller.logger.Info("header ip:" + headerIp + "customer header:" + headerIp2)

	ctx.JSON(http.StatusOK, controller.testUsecase.GetTestById(id))
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} Test Post
// @Param request body domain.TestModel true "Request"
// @Router /test [post]
func (controller *TestController) postTest(ctx *gin.Context) {
	var request domain.TestModel

	err := ctx.ShouldBind(&request)
	if err != nil {
		log.Println("postTest bind request Error!", err)

		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad Request!!!"})
		return
	}

	ctx.JSON(http.StatusOK, controller.testUsecase.PostTest(request))
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Patch
// @Param request body domain.TestModel true "Request"
// @Router /test/{id} [patch]
func (controller *TestController) patchTest(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var request domain.TestModel

	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad Request!!!"})
		return
	}

	ctx.JSON(http.StatusOK, controller.testUsecase.PatchTest(id, request))
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Delete
// @Router /test/{id} [delete]
func (controller *TestController) deleteTest(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, controller.testUsecase.DeleteTest(id))
}
