package main

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Authentication struct {
	Secret       string
	RedisClient  *redis.Client
	RedisContext context.Context
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	ATExpires    time.Time
	RTExpires    time.Time
}

type UserDetails struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type AccessDetails struct {
	Role   string
	UserId int
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type MyCustomClaims struct {
	Name string
	Uuid string
	Role string
	jwt.RegisteredClaims
}
