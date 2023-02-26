package authentication

import "github.com/jmoiron/sqlx"

type UserService struct {
	db *sqlx.DB
}

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstname" db:"firstname"`
	LastName   string `json:"lastname" db:"lastname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Phone      string `json:"phone" db:"phone"`
	Email      string `json:"email" db:"email"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}

type UserMethods interface {
	Get(userId int) (*User, error)
	Delete()
	Update()
}

func NewUserService(db *sqlx.DB) UserMethods {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Get(userId int) (*User, error) {
	var user *User

	return user, nil
}

func (u *UserService) Delete() {

}

func (u *UserService) Update() {

}
