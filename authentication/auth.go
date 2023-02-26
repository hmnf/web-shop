package authentication

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	JWTSecret string
	db        *sqlx.DB
	user      UserMethods
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

type AuthMethods interface {
	Login(w http.ResponseWriter, email string, password string) (*User, error)
	Registration()
	ExtractAccessMetaData()
	ExtractRefreshMetaData()
	ValidToken()
	GetTokenFromCookie()
	CreateToken(userId int) (*TokenDetails, error)
}

func NewAuthenticationService(db *sqlx.DB, jwtSecret string, user UserMethods) AuthMethods {
	return &Auth{
		JWTSecret: jwtSecret,
		db:        db,
		user:      user,
	}
}

func (auth *Auth) Login(w http.ResponseWriter, email string, password string) (*User, error) {
	var userId int
	var err error
	var td *TokenDetails

	// auth.db.Get(&userId, "SELECT id FROM users WHERE email=$1 AND password=$2", email, password)

	if userId > 0 {
		td, err = auth.CreateToken(userId)

		if err != nil {
			return nil, err
		}

		ATCookie := &http.Cookie{
			Name:     "access_token",
			Value:    td.AccessToken,
			Expires:  td.ATExpires,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		}

		RTCookie := &http.Cookie{
			Name:     "refresh_token",
			Value:    td.RefreshToken,
			Expires:  td.RTExpires,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		}

		http.SetCookie(w, ATCookie)
		http.SetCookie(w, RTCookie)

		user, err := auth.user.Get(userId)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, errors.New("wrong email or password")

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

func (auth *Auth) CreateToken(userId int) (*TokenDetails, error) {
	var err error

	atExpires := time.Now().Add(time.Minute * 15)
	rtExpires := time.Now().Add(time.Hour * 24 * 7)

	td := &TokenDetails{
		AccessUuid:  uuid.NewString(),
		RefreshUuid: uuid.NewString(),
		ATExpires:   atExpires,
		RTExpires:   rtExpires,
	}

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["ATExpires"] = td.ATExpires

	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userId
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["RTExpires"] = td.RTExpires

	clearToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = clearToken.SignedString([]byte(auth.JWTSecret))

	if err != nil {
		return nil, err
	}

	clearToken = jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = clearToken.SignedString([]byte(auth.JWTSecret))

	if err != nil {
		return nil, err
	}

	// auth.db.MustExec("INSERT INTO tokens (id, user_id, expires) VALUES ($1, $2, $3)", td.AccessUuid, userId, atExpires)
	// auth.db.MustExec("INSERT INTO tokens (id, user_id, expires) VALUES ($1, $2, $3)", td.RefreshUuid, userId, rtExpires)

	return td, nil
}
