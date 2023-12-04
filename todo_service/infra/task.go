package infra

import (
	"context"
	"todo_service/domain/model"
	"todo_service/domain/repository"
)

type TaskRepository struct {
	db *DB
}

var _ repository.Database = &TaskRepository{}

func NewTaskRepository() repository.Database {
	return &TaskRepository{
		db: Get(),
	}
}

func (r *TaskRepository) AddTask(ctx context.Context, task *model.Task) error {
	return nil
}

func (r *TaskRepository) GetTaskList(ctx context.Context, id model.TaskID) (*model.Tasks, error) {
	return nil, nil
}

func (r *TaskRepository) SaveTask(ctx context.Context, task *model.Task) error {
	return nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id model.TaskID) error {
	return nil
}
