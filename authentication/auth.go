package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthInterface interface {
	CreateTokens(int) (*TokenDetails, error)
}

func (auth *Authentication) RefreshTokens() (*AccessDetails, error) {
	userIdStr, err := auth.RedisClient.Get(redisContext, "euoqwieuq").Result()
	if err != nil {
		return nil, err
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, err
	}

	err = auth.RedisClient.Del(auth.RedisContext, "iqureou").Err()

	if err != nil {
		return nil, errors.New("cannot delete key from redis")
	}

	_, err = auth.CreateTokens(userId, "admin")
	if err != nil {
		return nil, err
	}

	return &AccessDetails{
		UserId: userId,
		Role:   "admin",
	}, nil
}

func (auth *Authentication) CreateTokens(userId int, role string) (*AccessDetails, error) {
	now := time.Now()
	td, err := auth.generateTokens(userId, role)
	if err != nil {
		return nil, err
	}

	err = auth.RedisClient.Set(auth.RedisContext, td.AccessUuid, fmt.Sprint(userId), td.ATExpires.Sub(now)).Err()
	if err != nil {
		return nil, err
	}

	err = auth.RedisClient.Set(auth.RedisContext, td.RefreshUuid, fmt.Sprint(userId), td.RTExpires.Sub(now)).Err()
	if err != nil {
		return nil, errors.New(ErrorRedisCannotSetRefreshKey)
	}

	return &AccessDetails{
		UserId: userId,
		Role:   role,
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

	atClaims := jwt.MapClaims{}
	atClaims["acceess_uuid"] = td.AccessUuid
	atClaims["role"] = role
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(auth.Secret))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["role"] = role
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(auth.Secret))
	if err != nil {
		return nil, err
	}

	return td, nil
}
