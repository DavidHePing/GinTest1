package repository

import (
	"GinTest1/domain"

	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (repo *CarRepository) GetCar(carId int) domain.Car {
	var car domain.Car
	repo.db.Find(&car)

	return car
}
