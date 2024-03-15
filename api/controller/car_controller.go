package controller

import (
	"GinTest1/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CarController struct {
	carUsecase domain.CarUseCase
}

func NewCarController(carUsecase domain.CarUseCase, logger *zap.Logger, router *gin.RouterGroup) {
	controller := &CarController{carUsecase: carUsecase}

	router.GET("/car/:id", controller.getCar)
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
