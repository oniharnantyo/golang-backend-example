package middleware

import (
	"github.com/go-redis/redis/v8"
	"net/http"
)

type (
	Middleware struct {
		RedisClient *redis.Client
	}
)

var middleware Middleware

func (m Middleware)Init()  {
	middleware = m
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorizatio
}
