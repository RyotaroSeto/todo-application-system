package service

import (
	"context"
	"user_service/domain/model"
)

type User interface {
	Register(ctx context.Context, user model.User) (model.UserID, error)
	Login(ctx context.Context, userName, password string) (string, error)
}
