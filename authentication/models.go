package main

import "time"

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
	TokenUuid string
	UserId    int
}
