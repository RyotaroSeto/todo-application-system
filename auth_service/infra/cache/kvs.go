package cache

import (
	"auth_service/domain/repository"
	"auth_service/infra/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type KvsRepository struct {
	redis *redis.Client
}

var _ repository.Cache = (*KvsRepository)(nil)

func NewRedisRepository(ctx context.Context, cfg *config.Config) *KvsRepository {
	return &KvsRepository{
		redis: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		}),
	}
}

func (r *KvsRepository) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (r *KvsRepository) Set(ctx context.Context, key string, value string) error {
	return nil
}
