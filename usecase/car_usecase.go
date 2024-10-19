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

type carUseCase struct {
	carRepository domain.CarRepository
}

func NewCarUsecase(repo domain.CarRepository) CarUseCase {
	return &carUseCase{carRepository: repo}
}

func (usecase carUseCase) GetAllCar() []*domain.Car {
	cars := usecase.carRepository.GetAllCar()
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Id < cars[j].Id
	})
	return cars
}

func (usecase carUseCase) GetCar(carId int) *domain.Car {
	car := usecase.carRepository.GetCar(carId)
	return car
}

// CreateCar implements domain.CarUseCase.
func (usecase *carUseCase) CreateCar(car *domain.Car) *domain.Car {
	return usecase.carRepository.CreateCar(car)
}

// UpdateCar implements domain.CarUseCase.
func (usecase *carUseCase) UpdateCar(carId int, car *domain.Car) *domain.Car {
	return usecase.carRepository.UpdateCar(carId, car)
}

// DeleteCar implements domain.CarUseCase.
func (usecase *carUseCase) DeleteCar(carId int) *domain.Car {
	return usecase.carRepository.DeleteCar(carId)
}
