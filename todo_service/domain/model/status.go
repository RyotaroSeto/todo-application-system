package model

import "time"

//go:generate go run github.com/dmarkham/enumer -type=TodoStatus -text -json -transform=snake-upper
type TodoStatus uint

const (
	TodoStatusUnspecified TodoStatus = iota
	TodoStatusDoing
	TodoStatusDone
)

// func

type Status struct {
	ID        TodoStatus `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;comment:削除日時"`
}

func (s *Status) TableName() string {
	return "status"
}
