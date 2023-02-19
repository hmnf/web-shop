package services

import "github.com/jmoiron/sqlx"

type UserService struct {
	db *sqlx.DB
}

type User struct {
	Id         int
	FirstName  string
	LastName   string
	Patronymic string
	Phone      string
	Email      string
	CreatedAt  string
	Cards      []Card
}

type UserMethods interface {
	Get()
	Update()
	Delete()
}

func NewUserService(db *sqlx.DB) UserMethods {
	return &UserService{
		db: db,
	}
}

func (user *UserService) Get() {

}

func (user *UserService) Update() {

}

func (user *UserService) Delete() {

}
