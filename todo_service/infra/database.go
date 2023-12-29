package infra

import (
	"database/sql"
	"pkg/errors"
	"time"
	"todo_service/domain/model"

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
			TablePrefix: "todo.",
		},
	}
)

func DBConnect(dsn string, maxConn, maxIdle int) error {
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
	if err = db.AutoMigrate(&model.Todo{}, &model.Status{}); err != nil {
		return errors.Errorf(errors.InternalServerError, "failed to migrate: %w", err)
	}

	if er := db.Where(&model.Status{ID: model.TodoStatus(1)}).First(&model.Status{}).Error; er != nil {
		if er == gorm.ErrRecordNotFound {
			return migrateStatusData(db)
		}
		return errors.Errorf(errors.InternalServerError, "failed to get status: %w", er)
	}
	return nil
}

func Get() *DB {
	return d
}

func migrateStatusData(db *gorm.DB) error {
	result := db.Create([]*model.Status{{ID: model.TodoStatus(1), Name: "Doing"}, {ID: model.TodoStatus(2), Name: "Done"}})
	return result.Error
}
