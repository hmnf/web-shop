package main

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
	Create() error
	Delete()
	Update()
}

func Create() error {
	return nil
}

func Delete() {

}

func Update() {

}
