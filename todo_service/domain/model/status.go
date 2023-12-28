package model

import "time"

type StatusID int64

//go:generate go run github.com/dmarkham/enumer -type=TodoStatus -text -json -transform=snake-upper
type TodoStatus uint

const (
	TodoStatusUnspecified TodoStatus = iota
	TodoStatusDoing
	TodoStatusDone
)

type Status struct {
	ID        StatusID   `gorm:"primaryKey"`
	Status    TodoStatus `gorm:"index;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;comment:削除日時"`
}
