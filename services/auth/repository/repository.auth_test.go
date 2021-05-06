package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-redis/redis/v8"
	redismock "github.com/go-redis/redismock/v8"
)

var (
	key = "key"
	val = "val"
)

func TestAuthRepository_Set(t *testing.T) {
	ctx := context.Background()

	exp := time.Duration(1)

	client, mockRedis := redismock.NewClientMock()

	t.Run("Success", func(t *testing.T) {
		mockRedis.ExpectSet(key, val, exp).SetVal("")
		a := NewAuthRepository(client)
		err := a.Set(ctx, key, val, exp)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRedis.ExpectSet(key, val, exp).SetErr(redis.ErrClosed)
		a := NewAuthRepository(client)
		err := a.Set(ctx, key, val, exp)
		assert.Error(t, err)
	})
}

func TestAuthRepository_Get(t *testing.T) {
	ctx := context.Background()

	client, mockRedis := redismock.NewClientMock()

	t.Run("Success", func(t *testing.T) {
		mockRedis.ExpectGet(key).SetVal(val)

		a := NewAuthRepository(client)
		res, err := a.Get(ctx, key)

		assert.NoError(t, err)
		assert.Equal(t, val, res)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRedis.ExpectGet(key).RedisNil()

		a := NewAuthRepository(client)
		res, err := a.Get(ctx, key)
		assert.Error(t, err)
		assert.NotEqual(t, val, res)
	})
}
