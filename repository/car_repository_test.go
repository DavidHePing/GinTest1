package repository

import (
	"GinTest1/domain"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestGetAllCar(t *testing.T) {
	mock, db, gormDB := getGormMock(t)
	defer db.Close()

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

	mock.ExpectQuery(`SELECT \* FROM "cars" WHERE deleted_at is null`).
		WillReturnRows(rows)
	repo := NewCarRepository(gormDB)
	car := repo.GetAllCar()

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, car)
	assert.Len(t, car, 2)
	assert.Equal(t, car, mockCars)
}

func TestGetCar(t *testing.T) {
	mock, db, gormDB := getGormMock(t)
	defer db.Close()

	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	rows := sqlmock.NewRows(([]string{"id", "name", "price", "type"})).
		AddRow(mockCar.Id, mockCar.Name, mockCar.Price, mockCar.Type)

	mock.ExpectQuery(`SELECT \* FROM "cars" WHERE "cars"."id" = \$?`).
		WithArgs(1).
		WillReturnRows(rows)
	repo := NewCarRepository(gormDB)
	car := repo.GetCar(1)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, car)
	assert.Equal(t, car, &mockCar)
}

func TestCreateCar(t *testing.T) {
	mock, db, gormDB := getGormMock(t)
	defer db.Close()

	// Set up the mock car object
	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	}

	// Create mock rows that match the expected query result
	rows := sqlmock.NewRows([]string{"id", "name", "price", "type"}).
		AddRow(mockCar.Id, mockCar.Name, mockCar.Price, mockCar.Type)

	// Expect a transaction to begin
	mock.ExpectBegin()

	// Use a regular expression to match the parameterized query
	mock.ExpectQuery(`INSERT INTO "cars" \("name","price","type"\) VALUES \(\$1,\$2,\$3\) RETURNING "id"`).
		WithArgs("Tesla", 100.0, "US").
		WillReturnRows(rows)

	// Expect the transaction to commit
	mock.ExpectCommit()

	// Initialize the repository and call CreateCar
	repo := NewCarRepository(gormDB)
	car := repo.CreateCar(&domain.Car{
		Name:  "Tesla",
		Price: 100,
		Type:  "US",
	})

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, car)
	assert.Equal(t, car, &mockCar)
}

func TestUpdateCar(t *testing.T) {
	mock, db, gormDB := getGormMock(t)
	defer db.Close()

	// Set up the mock car object
	mockCar := domain.Car{
		Id:    1,
		Name:  "Tesla1",
		Price: 100,
		Type:  "US",
	}

	// Expect a transaction to begin
	mock.ExpectBegin()

	// Use a regular expression to match the parameterized query
	mock.ExpectExec(`UPDATE "cars" SET "name"=\$1,"price"=\$2,"type"=\$3 WHERE "id" = \$4`).
		WithArgs("Tesla1", 100.0, "US", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect the transaction to commit
	mock.ExpectCommit()

	// Initialize the repository and call UpdateCar
	repo := NewCarRepository(gormDB)
	car := repo.UpdateCar(1, &domain.Car{
		Name:  "Tesla1",
		Price: 100,
		Type:  "US",
	})

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, car)
	assert.Equal(t, car, &mockCar)
}

func getGormMock(t *testing.T) (mock sqlmock.Sqlmock, db *sql.DB, gormDB *gorm.DB) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Failed to open gorm DB: %v", err)
	}

	return
}
