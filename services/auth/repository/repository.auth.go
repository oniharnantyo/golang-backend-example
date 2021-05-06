package repository

import (
	"context"
	"time"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/go-redis/redis/v8"
)

type authRepository struct {
	redisClient *redis.Client
}

func (a authRepository) Get(ctx context.Context, id string) (string, error) {
	res, err := a.redisClient.Get(ctx, id).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (a authRepository) Set(ctx context.Context, key, value string, expire time.Duration) error {
	err := a.redisClient.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}

	return nil
}

func NewAuthRepository(client *redis.Client) domain.AuthRepository {
	return &authRepository{redisClient: client}
}
