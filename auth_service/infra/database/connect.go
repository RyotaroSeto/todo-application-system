package database

import (
	"pkg/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var d *DB

func Connect(dsn string, maxConn, maxIdle int) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return errors.Errorf(errors.InternalServerError, "failed to get sql.DB: %w", err)
	}
	if err = sqlDB.Ping(); err == nil {
		return errors.Errorf(errors.InternalServerError, "failed to ping: %w", err)
	}

	if maxConn > 0 {
		sqlDB.SetMaxOpenConns(maxConn)
	}
	if maxIdle > 0 {
		sqlDB.SetMaxIdleConns(maxIdle)
	}

	d = &DB{db}

	return nil
}

func Get() *DB {
	return d
}
