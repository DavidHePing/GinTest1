package domain

type Car struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Price     float64
	IsDeleted bool
}

func (Car) TableName() string {
	return "test_cars"
}

type CarRepository interface {
	GetCar(carId int) Car
}

type CarUseCase interface {
	GetCar(carId int) Car
}
