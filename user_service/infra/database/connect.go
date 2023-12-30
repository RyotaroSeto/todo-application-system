package database

import (
	"database/sql"
	"pkg/errors"
	"time"
	"user_service/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const retryLimit = 10

type DB struct {
	*gorm.DB
}

var (
	d    *DB
	opts = gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "user.",
		},
	}
)

func Connect(dsn string, maxConn, maxIdle int) error {
	db, err := gorm.Open(postgres.Open(dsn), &opts)
	if err != nil {
		return err
	}

	var sqlDB *sql.DB
	for i := 0; i < retryLimit; i++ {
		time.After(1 * time.Second)
		sqlDB, err = db.DB()
		if err != nil {
			return errors.Errorf(errors.InternalServerError, "failed to get sql.DB: %w", err)
		}

		if err = sqlDB.Ping(); err == nil {
			break
		}
	}

	if maxConn > 0 {
		sqlDB.SetMaxOpenConns(maxConn)
	}
	if maxIdle > 0 {
		sqlDB.SetMaxIdleConns(maxIdle)
	}

	d = &DB{db}

	return db.AutoMigrate(&model.User{})
}

func Get() *DB {
	return d
}
