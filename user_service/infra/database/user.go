package database

import (
	"context"
	"user_service/domain/model"
	"user_service/domain/repository"
)

type DatabaseRepository struct {
	db *DB
}

var _ repository.Database = &DatabaseRepository{}

func NewTodoRepository() repository.Database {
	return &DatabaseRepository{
		db: Get(),
	}
}

func (r *DatabaseRepository) RegisterUser(ctx context.Context, user model.User) (model.UserID, error) {
	// result := r.db.Create(user)

	return 0, nil
	// return user.ID, result.Error
}

func (r *DatabaseRepository) GetUser(ctx context.Context, userName string) (*model.User, error) {
	user := &model.User{}
	// if err := r.db.Find(user).Preload(clause.Associations).Error; err != nil {
	// 	return nil, err
	// }
	return user, nil
}
