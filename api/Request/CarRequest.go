package request

import (
	enum "GinTest1/api/Enum"
	"GinTest1/domain"
)

type CarRequest struct {
	Name  string       `form:"Name" json:"Name" binding:"required"`
	Price float64      `form:"Price" json:"Price" binding:"required"`
	Type  enum.CarType `form:"Type" json:"Type" binding:"required"`
}

func (req CarRequest) ToCar() *domain.Car {
	return &domain.Car{Name: req.Name, Price: req.Price, Type: req.Type.GetDesciption()}
}
