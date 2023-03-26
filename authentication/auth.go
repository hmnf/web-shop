package main

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	JWTSecret string
	db        *sqlx.DB
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	ATExpires    time.Time
	RTExpires    time.Time
}

type AccessDetails struct {
	TokenUuid string
	UserId    int
}

type AuthInterface interface {
	CreateTokens(int) (*TokenDetails, error)
}

func CreateTokens(userId int) (*TokenDetails, error) {
	atExpires := time.Now().Add(time.Minute * 15)
	rtExpires := time.Now().Add(time.Hour * 7 * 24)

	td := &TokenDetails{
		AccessUuid:  uuid.NewString(),
		RefreshUuid: uuid.NewString(),
		ATExpires:   atExpires,
		RTExpires:   rtExpires,
	}

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["accessUuid"] = td.AccessUuid
	atClaims["expires"] = td.ATExpires

	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userId
	rtClaims["refreshUuid"] = td.RefreshUuid
	rtClaims["expires"] = td.RTExpires
}
