package services

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Auth struct {
	JWTSecret string
	user      UserMethods
	db        *sqlx.DB
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	ATExpires    time.Time
	AtExpires    time.Time
}

type AccessDetails struct {
	TokenUuid string
	UserId    int
}

type AuthMethods interface {
	Login()
	Registration()
	ExtractAccessMetaData()
	ExtractRefreshMetaData()
	ValidToken()
	GetTokenFromCookie()
	CreateToken()
}

func NewAuthorizationService(db *sqlx.DB, jwtSecret string, user UserMethods) AuthMethods {
	return &Auth{
		JWTSecret: jwtSecret,
		user:      user,
		db:        db,
	}
}

func (auth *Auth) Login() {

}

func (auth *Auth) Registration() {

}

func (auth *Auth) ExtractAccessMetaData() {

}

func (auth *Auth) ExtractRefreshMetaData() {

}

func (auth *Auth) ValidToken() {

}

func (auth *Auth) GetTokenFromCookie() {

}

func (auth *Auth) CreateToken() {

}
