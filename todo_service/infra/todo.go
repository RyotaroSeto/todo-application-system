package infra

import (
	"context"
	"todo_service/domain/model"
	"todo_service/domain/repository"

	"gorm.io/gorm/clause"
)

type TodoRepository struct {
	db *DB
}

var _ repository.Database = &TodoRepository{}

func NewTodoRepository() repository.Database {
	return &TodoRepository{
		db: Get(),
	}
}

func (r *TodoRepository) AddTodo(ctx context.Context, todo *model.Todo) (uint64, error) {
	result := r.db.Create(todo)
	return todo.ID, result.Error
}

func (r *TodoRepository) GetTodoList(ctx context.Context) (*model.Todos, error) {
	todo := &model.Todos{}
	if err := r.db.Find(todo).Preload(clause.Associations).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) UpdateTodoStatus(ctx context.Context, id uint64, statusID model.TodoStatus) error {
	return r.db.Model(&model.Todo{}).Where(model.Todo{ID: uint64(id)}).Updates(&model.Todo{TodoStatus: statusID}).Error
}

func (r *TodoRepository) UpdateTodoTitle(ctx context.Context, id uint64, title model.TodoTitle) error {
	return r.db.Model(&model.Todo{}).Where(model.Todo{ID: uint64(id)}).Updates(&model.Todo{Title: title}).Error
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id uint64) error {
	return r.db.Where(model.Todo{ID: uint64(id)}).Delete(&model.Todo{}).Error
}
