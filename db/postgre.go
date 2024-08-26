package db

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgre struct {
}

func (Postgre) GetGormDb(connectString string, fileLogger *zap.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
		//show all sql
		Logger: logger.Default.LogMode(logger.Info),
		//not do only show sql
		// DryRun: true,
		//skip transaction, performance 30% up!
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fileLogger.Error(err.Error())
		panic("Db connect failed!")
	}

	// Get the generic database object sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		fileLogger.Error(err.Error())
		panic("Failed to get sql.DB from gorm.DB")
	}

	// Configure the connection pool
	sqlDB.SetMaxIdleConns(10)           // Set the maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)          // Set the maximum number of open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Set the maximum connection lifetime

	return db
}
