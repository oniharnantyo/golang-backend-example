package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type authUseCase struct {
	authRepository                domain.AuthRepository
	AccessSecret                  string
	AccessSecretExpireAfterMinute int
	RefreshSecret                 string
	RefreshSecretExpireAfterDay   int
}

func (a authUseCase) CreateAuth(ctx context.Context, account domain.Account) (domain.Auth, error) {
	tokenData := domain.Auth{
		AccessUuid:           uuid.New().String(),
		RefreshUuid:          uuid.New().String(),
		AccessTokenExpireAt:  time.Now().Add(time.Minute * time.Duration(a.AccessSecretExpireAfterMinute)).Unix(),
		RefreshTokenExpireAt: time.Now().Add(time.Hour * 24 * time.Duration(a.RefreshSecretExpireAfterDay)).Unix(),
	}

	accessClaims := domain.AccessClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(a.AccessSecretExpireAfterMinute)).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		},
		Account: &account,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := at.SignedString([]byte(a.AccessSecret))
	if err != nil {
		return domain.Auth{}, errors.Wrap(err, "authUseCase/CreateAuth/SignedString")
	}
	tokenData.AccessToken = accessToken

	//Refresh Claims
	refreshClaims := domain.RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(a.RefreshSecretExpireAfterDay)).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		},
		Account: &account,
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := rt.SignedString([]byte(a.RefreshSecret))
	if err != nil {
		return domain.Auth{}, errors.Wrap(err, "authUseCase/CreateAuth/SignedString")
	}

	tokenData.RefreshToken = refreshToken

	//Save token to Redis
	err = a.authRepository.Set(ctx, tokenData.AccessUuid, strconv.Itoa(account.AccountNumber), time.Duration(a.AccessSecretExpireAfterMinute)*time.Minute)
	if err != nil {
		return domain.Auth{}, errors.Wrap(err, "authUseCase/CreateAuth/SetAccessID")
	}

	err = a.authRepository.Set(ctx, tokenData.RefreshUuid, strconv.Itoa(account.AccountNumber), time.Duration(a.RefreshSecretExpireAfterDay)*24*time.Hour)
	if err != nil {
		return domain.Auth{}, errors.Wrap(err, "authUseCase/CreateAuth/SetRefreshID")
	}

	return tokenData, nil
}

func NewAuthUseCase(
	a domain.AuthRepository,
	accessSecret string,
	accessSecretExpireAfterMinute int,
	refreshSecret string,
	refreshSecretExpireAfterDay int,
) domain.AuthUseCase {
	return &authUseCase{
		authRepository:                a,
		AccessSecret:                  accessSecret,
		AccessSecretExpireAfterMinute: accessSecretExpireAfterMinute,
		RefreshSecret:                 refreshSecret,
		RefreshSecretExpireAfterDay:   refreshSecretExpireAfterDay,
	}
}
