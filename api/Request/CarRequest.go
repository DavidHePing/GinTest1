package request

import "GinTest1/domain"

type CarRequest struct {
	Name  string
	Price float64
	Type  string
}

func (req CarRequest) ToCar() *domain.Car {
	return &domain.Car{Name: req.Name, Price: req.Price, Type: req.Type}
}
