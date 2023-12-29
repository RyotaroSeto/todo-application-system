package service

import (
	"context"
	"todo_service/domain/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -destination=interface_mock.go -package=service
type Todo interface {
	Add(ctx context.Context, todo *model.Todo) error
	GetList(ctx context.Context) (*model.Todo, error)
	UpdateStatus(ctx context.Context, id uint64, statusID model.TodoStatus) error
	UpdateTitle(ctx context.Context, id uint64, title model.TodoTitle) error
	Delete(ctx context.Context, id uint64) error
}
