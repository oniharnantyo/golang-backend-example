package domain

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	Auth struct {
		AccessToken          string `json:"access_token"`
		RefreshToken         string `json:"refresh_token"`
		AccessUuid           string `json:"access_uuid"`
		RefreshUuid          string `json:"refresh_uuid"`
		AccessTokenExpireAt  int64  `json:"access_token_expire_at"`
		RefreshTokenExpireAt int64  `json:"refresh_token_expire_at"`
	}

	AccessClaims struct {
		jwt.StandardClaims
		AccessUUID string   `json:"access_uuid"`
		Account    *Account `json:"account"`
	}

	RefreshClaims struct {
		jwt.StandardClaims
		RefreshUUID string   `json:"refresh_uuid"`
		Account     *Account `json:"account"`
	}
)

type (
	AuthUseCase interface {
		CreateAuth(ctx context.Context, account Account) (Auth, error)
	}

	AuthRepository interface {
		Get(ctx context.Context, id string) (string, error)
		Set(ctx context.Context, key, value string, expire time.Duration) error
	}
)
