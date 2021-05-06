package util

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type (
	Redis struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
)

func (r Redis) Init(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(`%s:%d`, r.Host, r.Port),
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
