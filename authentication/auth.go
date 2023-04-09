package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthInterface interface {
	CreateTokens(int) (*TokenDetails, error)
}

func (auth *Authentication) RefreshTokens(RT string) (*Tokens, error) {
	parsedToken, err := auth.validateToken(RT)

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*MyCustomClaims); ok && parsedToken.Valid {
		userIdStr, err := auth.RedisClient.Get(redisContext, claims.Uuid).Result()
		if err != nil {
			return nil, errors.New(ErrorRedisTokenIsInvalid)
		}

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return nil, err
		}

		err = auth.RedisClient.Del(auth.RedisContext, claims.Uuid).Err()

		if err != nil {
			return nil, errors.New(ErrorRedisCannotDeleteKey)
		}

		return auth.CreateTokens(userId, "admin")
	} else {
		return nil, err
	}
}

func (auth *Authentication) validateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(auth.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}

func (auth *Authentication) CreateTokens(userId int, role string) (*Tokens, error) {
	td, err := auth.generateTokens(userId, role)
	if err != nil {
		return nil, err
	}

	err = auth.RedisClient.Set(redisContext, td.AccessUuid, fmt.Sprint(userId), time.Until(td.ATExpires)).Err()
	if err != nil {
		return nil, errors.New(ErrorRedisCannotSetRefreshKey)
	}

	err = auth.RedisClient.Set(redisContext, td.RefreshUuid, fmt.Sprint(userId), time.Until(td.RTExpires)).Err()
	if err != nil {
		return nil, errors.New(ErrorRedisCannotSetRefreshKey)
	}

	return &Tokens{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}, nil
}

func (auth *Authentication) generateTokens(userId int, role string) (*TokenDetails, error) {
	atExpires := time.Now().Add(time.Minute * 15)
	rtExpires := time.Now().Add(time.Hour * 7 * 24)

	td := &TokenDetails{
		AccessUuid:  uuid.NewString(),
		RefreshUuid: uuid.NewString(),
		ATExpires:   atExpires,
		RTExpires:   rtExpires,
	}

	var err error

	atClaims := MyCustomClaims{
		"ars",
		td.AccessUuid,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(atExpires),
			Issuer:    "ars",
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(auth.Secret))
	if err != nil {
		return nil, err
	}

	rtClaims := MyCustomClaims{
		"ars",
		td.RefreshUuid,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(rtExpires),
			Issuer:    "ars",
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(auth.Secret))
	if err != nil {
		return nil, err
	}

	return td, nil
}
