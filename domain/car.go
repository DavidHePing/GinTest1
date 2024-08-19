package domain

type Car struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Price float64
	Type  string
}

type CarRepository interface {
	GetAllCar() []*Car
	GetCar(carId int) *Car
	CreateCar(car *Car) bool
	UpdateCar(carId int, car *Car) bool
	DeleteCar(carId int) *Car
}

type CarUseCase interface {
	GetAllCar() []*Car
	GetCar(carId int) *Car
	CreateCar(car *Car) bool
	UpdateCar(carId int, car *Car) bool
	DeleteCar(carId int) *Car
}
