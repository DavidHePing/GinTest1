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
	CreateCar(car *Car) *Car
	UpdateCar(carId int, car *Car) *Car
	DeleteCar(carId int) *Car
}

type CarUseCase interface {
	GetAllCar() []*Car
	GetCar(carId int) *Car
	CreateCar(car *Car) *Car
	UpdateCar(carId int, car *Car) *Car
	DeleteCar(carId int) *Car
}
