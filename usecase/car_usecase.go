package usecase

import (
	"GinTest1/domain"
)

type carUsecase struct {
	carRepository domain.CarRepository
}

func NewCarUsecase(repo domain.CarRepository) *carUsecase {
	return &carUsecase{carRepository: repo}
}

func (usecase carUsecase) GetCar(carId int) domain.Car {
	car := usecase.carRepository.GetCar(carId)
	return car
}
