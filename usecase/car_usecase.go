package usecase

import (
	"GinTest1/domain"
	"GinTest1/repository"
)

type carUsecase struct {
	carRepository repository.CarRepository
}

func NewCarUsecase(repo repository.CarRepository) *carUsecase {
	return &carUsecase{carRepository: repo}
}

func (usecase carUsecase) GetCar(carId int) domain.Car {
	car := usecase.carRepository.GetCar(carId)
	return car
}
