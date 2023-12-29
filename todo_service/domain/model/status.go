package model

import "time"

//go:generate go run github.com/dmarkham/enumer -type=TodoStatus -text -json -transform=snake-upper
type TodoStatus uint

const (
	TodoStatusUnspecified TodoStatus = iota
	TodoStatusDoing
	TodoStatusDone
)

// type TodoName string

// const (
// 	TodoNameUnspecified TodoStatus = iota
// 	TodoNameDoing
// 	TodoNameDone
// )

// func (t TodoName) String() string {
// 	return string(t)
// }

type Status struct {
	ID        TodoStatus `gorm:"primaryKey"`
	Name      string     `gorm:"not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;comment:削除日時"`
}

func (s *Status) TableName() string {
	return "status"
}
