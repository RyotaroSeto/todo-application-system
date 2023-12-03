package repository

import (
	"context"
	"todo_service/domain/model"
)

type TodoRepository interface {
	AddTask(ctx context.Context, task *model.Task) error
	GetTaskList(ctx context.Context, id model.TaskID) (*model.Tasks, error)
	SaveTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id model.TaskID) error
}
