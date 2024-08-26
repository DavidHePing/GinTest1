package usecase

import (
	"GinTest1/domain"
	"GinTest1/domain/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCar(t *testing.T) {
	carUsecase, carRepo := getCarUseCase(t)

	mockCars := []*domain.Car{
		{
			Id:    1,
			Name:  "Tesla",
			Price: 100,
			Type:  "US",
		},
		{
			Id:    2,
			Name:  "Honda",
			Price: 60,
			Type:  "Japan",
		},
	}

	carRepo.EXPECT().GetAllCar().Return(mockCars)
	cars := carUsecase.GetAllCar()

	assert.NotNil(t, cars)
	assert.Len(t, cars, 2)
	assert.Equal(t, mockCars, cars)
}

func TestGetCar(t *testing.T) {
	carUsecase, carRepo := getCarUseCase(t)

	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	carRepo.EXPECT().GetCar(1).Return(&mockCar)
	car := carUsecase.GetCar(1)

	assert.NotNil(t, car)
	assert.Equal(t, &mockCar, car)
}

func TestCreateCar(t *testing.T) {
	carUsecase, carRepo := getCarUseCase(t)

	newCar := domain.Car{
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	carRepo.EXPECT().CreateCar(&newCar).Return(&mockCar)
	car := carUsecase.CreateCar(&newCar)

	assert.NotNil(t, car)
	assert.Equal(t, &mockCar, car)
}

func TestUpdateCar(t *testing.T) {
	carUsecase, carRepo := getCarUseCase(t)

	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	carRepo.EXPECT().UpdateCar(1, &mockCar).Return(&mockCar)
	car := carUsecase.UpdateCar(1, &mockCar)

	assert.NotNil(t, car)
	assert.Equal(t, &mockCar, car)
}

func TestDeleteCar(t *testing.T) {
	carUsecase, carRepo := getCarUseCase(t)

	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	carRepo.EXPECT().DeleteCar(1).Return(&mockCar)
	car := carUsecase.DeleteCar(1)

	assert.NotNil(t, car)
	assert.Equal(t, &mockCar, car)
}

func getCarUseCase(t *testing.T) (domain.CarUseCase, *mocks.MockCarRepository) {
	ctrl := gomock.NewController(t)

	carRepo := mocks.NewMockCarRepository(ctrl)
	return NewCarUsecase(carRepo), carRepo
}
