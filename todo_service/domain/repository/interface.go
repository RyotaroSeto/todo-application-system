package repository

import (
	"context"
	"todo_service/domain/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -package=repository -destination=interface_mock.go
type Database interface {
	AddTask(ctx context.Context, task *model.Task) error
	GetTaskList(ctx context.Context, id model.TaskID) (*model.Tasks, error)
	SaveTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id model.TaskID) error
}
