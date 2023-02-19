package Auth

import "time"

type Auth struct {
	JWTSecret string
	user      string
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

func NewAuthorizationService(jwtSecret string) AuthMethods {
	return &Auth{
		JWTSecret: jwtSecret,
	}
}

func (auth *Auth) Login() {

}
