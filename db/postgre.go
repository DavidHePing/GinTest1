package db

import (
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

	return db
}
