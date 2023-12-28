package model

import "time"

type Todo struct {
	ID        uint64    `gorm:"primaryKey"`
	Title     TodoTitle `gorm:"not null;"`
	StatusID  `gorm:"index;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;comment:削除日時"`
}

type TodoTitle string

type Todos []Todo
