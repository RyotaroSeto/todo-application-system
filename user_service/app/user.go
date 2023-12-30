package app

import (
	"context"
	"user_service/domain/model"
	"user_service/domain/repository"
)

type UserService struct {
	repo repository.Database
}

func NewUserService(repo repository.Database) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, user model.User) (model.UserID, error) {
	// パスワード処理など
	return s.repo.RegisterUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, userName, password string) (string, error) {
	// user, err := s.repo.GetUser(ctx, userName)
	// if err != nil {
	// 	return "", err
	// }

	// パスワード処理など
	return "", nil
}
