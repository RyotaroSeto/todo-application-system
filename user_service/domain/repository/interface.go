package repository

import (
	"context"
	"user_service/domain/model"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type Database interface {
	RegisterUser(ctx context.Context, user model.User) (model.UserID, error)
	GetUser(ctx context.Context, userName string) (*model.User, error)
}
