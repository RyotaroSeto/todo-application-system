package model

import (
	"strings"
	"time"
)

//go:generate go run github.com/dmarkham/enumer -type=TodoStatus -text -json -transform=snake-upper
type TodoStatus uint

const todoStatusNamePrefix = "TODO_STATUS_"

const (
	TodoStatusUnspecified TodoStatus = iota
	TodoStatusDoing
	TodoStatusDone
)

func (s TodoStatus) ShortString() string {
	return strings.ReplaceAll(s.String(), todoStatusNamePrefix, "")
}

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
