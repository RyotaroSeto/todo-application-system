package repository

import (
	"context"
	"todo_service/domain/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -destination=interface_mock.go -package=repository
type Database interface {
	AddTodo(ctx context.Context, todo *model.Todo) error
	GetTodoList(ctx context.Context, id uint64) (*model.Todo, error)
	UpdateTodoStatus(ctx context.Context, id uint64, status model.TodoStatus) error
	UpdateTodoTitle(ctx context.Context, id uint64, title model.TodoTitle) error
	DeleteTodo(ctx context.Context, id uint64) error
}
