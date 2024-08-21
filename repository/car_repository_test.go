package repository

import (
	"GinTest1/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestGetAllCar(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm DB: %v", err)
	}

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

	rows := sqlmock.NewRows(([]string{"id", "name", "price", "type"})).
		AddRow(mockCars[0].Id, mockCars[0].Name, mockCars[0].Price, mockCars[0].Type).
		AddRow(mockCars[1].Id, mockCars[1].Name, mockCars[1].Price, mockCars[1].Type)

	mock.ExpectQuery("SELECT \\* FROM \"cars\" WHERE deleted_at is null").
		WillReturnRows(rows)
	repo := NewCarRepository(gormDB)
	car := repo.GetAllCar()

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, car)
	assert.Len(t, car, 2)
	assert.Equal(t, car, mockCars)
}

func TestGetCar(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL logging
	})
	if err != nil {
		t.Fatalf("Failed to open gorm DB: %v", err)
	}

	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	rows := sqlmock.NewRows(([]string{"id", "name", "price", "type"})).
		AddRow(mockCar.Id, mockCar.Name, mockCar.Price, mockCar.Type)

	mock.ExpectQuery("SELECT \\* FROM \"cars\" WHERE \"cars\".\"id\" = \\$?").
		WithArgs(1).
		WillReturnRows(rows)
	repo := NewCarRepository(gormDB)
	car := repo.GetCar(1)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, car)
	assert.Equal(t, car, &mockCar)
}
