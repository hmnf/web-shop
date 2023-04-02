package main

import (
	"context"
	"time"

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
