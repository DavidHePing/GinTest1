package controller

import (
	request "GinTest1/api/Request"
	"GinTest1/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

type CarController struct {
	carUsecase domain.CarUseCase
	cache      *cache.Cache
}

func NewCarController(carUsecase domain.CarUseCase, logger *zap.Logger,
	router *gin.RouterGroup, cache *cache.Cache) {

	controller := &CarController{carUsecase: carUsecase, cache: cache}

	router.GET("/car", controller.getAllCar)
	router.GET("/car/:id", controller.getCar)
	router.POST("/car", controller.createCar)
	router.PATCH("/car/:id", controller.updateCar)
	router.DELETE("/car/:id", controller.deleteCar)
}

// @Schemes
// @Tags car
// @Accept json
// @Produce json
// @Success 200 {string} Get all car
// @Router /car [get]
func (controller *CarController) getAllCar(ctx *gin.Context) {
	if x, found := controller.cache.Get("getAllCar"); found {
		ctx.JSON(http.StatusOK, x)
		return
	}

	cars := controller.carUsecase.GetAllCar()
	controller.cache.Set("getAllCar", cars, 10*time.Second)
	ctx.JSON(http.StatusOK, cars)
}

// @Schemes
// @Tags car
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Get car by id
// @Router /car/{id} [get]
func (controller *CarController) getCar(ctx *gin.Context) {
	carId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad Request!!!"})
		return
	}

	ctx.JSON(http.StatusOK, controller.carUsecase.GetCar(carId))
}

// @Schemes
// @Tags car
// @Accept json
// @Produce json
// @Success 200 {string} Create car
// @Param request body request.CarRequest true "Request"
// @Router /car [post]
func (controller *CarController) createCar(ctx *gin.Context) {
	var request request.CarRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad Request!!!"})
		return
	}

	ctx.JSON(http.StatusOK, controller.carUsecase.CreateCar(request.ToCar()))
}

// @Schemes
// @Tags car
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Update car by id
// @Param request body request.CarRequest true "Request"
// @Router /car/{id} [patch]
func (controller *CarController) updateCar(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var request request.CarRequest

	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad Request!!!"})
		return
	}

	ctx.JSON(http.StatusOK, controller.carUsecase.UpdateCar(id, request.ToCar()))
}

// @Schemes
// @Tags car
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Delete car by id
// @Router /car/{id} [delete]
func (controller *CarController) deleteCar(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, controller.carUsecase.DeleteCar(id))
}
