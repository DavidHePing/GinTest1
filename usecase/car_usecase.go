package usecase

import (
	"GinTest1/domain"
	"sort"
)

type CarUseCase interface {
	GetAllCar() []*domain.Car
	GetCar(carId int) *domain.Car
	CreateCar(car *domain.Car) *domain.Car
	UpdateCar(carId int, car *domain.Car) *domain.Car
	DeleteCar(carId int) *domain.Car
}

type carUsecase struct {
	carRepository domain.CarRepository
}

func NewCarUsecase(repo domain.CarRepository) *carUsecase {
	return &carUsecase{carRepository: repo}
}

func (usecase carUsecase) GetAllCar() []*domain.Car {
	cars := usecase.carRepository.GetAllCar()
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Id < cars[j].Id
	})
	return cars
}

func (usecase carUsecase) GetCar(carId int) *domain.Car {
	car := usecase.carRepository.GetCar(carId)
	return car
}

// CreateCar implements domain.CarUseCase.
func (usecase *carUsecase) CreateCar(car *domain.Car) *domain.Car {
	return usecase.carRepository.CreateCar(car)
}

// UpdateCar implements domain.CarUseCase.
func (usecase *carUsecase) UpdateCar(carId int, car *domain.Car) *domain.Car {
	return usecase.carRepository.UpdateCar(carId, car)
}

// DeleteCar implements domain.CarUseCase.
func (usecase *carUsecase) DeleteCar(carId int) *domain.Car {
	return usecase.carRepository.DeleteCar(carId)
}
