package repository

import (
	"GinTest1/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	mockCars := []domain.Car{
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
	assert.Len(t, car, 2)
}
