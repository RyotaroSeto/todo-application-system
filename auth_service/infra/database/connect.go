package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var d *DB

func Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	d = &DB{db}

	return nil
}
func Get() *DB {
	return d
}
