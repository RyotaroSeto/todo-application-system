package app

import (
	"context"
	"todo_service/domain/model"
	"todo_service/domain/repository"
	"todo_service/domain/service"
)

type TodoService struct {
	repo repository.Database
}

var _ service.Todo = &TodoService{}

func NewTodoService(repo repository.Database) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) Add(ctx context.Context, todo *model.Todo) (uint64, error) {
	return s.repo.AddTodo(ctx, todo)
}

func (s *TodoService) GetList(ctx context.Context) (*model.Todos, error) {
	return s.repo.GetTodoList(ctx)
}

func (s *TodoService) UpdateStatus(ctx context.Context, id uint64, statusID model.TodoStatus) error {
	return s.repo.UpdateTodoStatus(ctx, id, statusID)
}

func (s *TodoService) UpdateTitle(ctx context.Context, id uint64, title model.TodoTitle) error {
	return s.repo.UpdateTodoTitle(ctx, id, title)
}

func (s *TodoService) Delete(ctx context.Context, id uint64) error {
	return s.repo.DeleteTodo(ctx, id)
}
