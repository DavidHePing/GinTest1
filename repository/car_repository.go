package repository

import (
	"GinTest1/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (repo *CarRepository) GetAllCar() []*domain.Car {
	var cars = []*domain.Car{}
	repo.db.Where("deleted_at is null").Find(&cars)

	return cars
}

func (repo *CarRepository) GetCar(carId int) *domain.Car {
	car := domain.Car{}
	repo.db.Find(&car, carId)

	return &car
}

// CreateCar implements domain.CarRepository.
func (repo *CarRepository) CreateCar(car *domain.Car) *domain.Car {
	repo.db.Omit("Id").Create(car)

	return car
}

// UpdateCar implements domain.CarRepository.
func (repo *CarRepository) UpdateCar(carId int, car *domain.Car) *domain.Car {
	car.Id = carId
	repo.db.Omit("Id").Save(car)
	return car
}

// DeleteCar implements domain.CarRepository.
func (repo *CarRepository) DeleteCar(carId int) *domain.Car {
	car := domain.Car{Id: carId}
	repo.db.Clauses(clause.Returning{}).Delete(&car)
	return &car
}
