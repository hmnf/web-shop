package main

type UserCreate struct {
	Firstname  string `json:"firstname" db:"firstname"`
	Lastname   string `json:"lastname" db:"lastname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Phone      string `json:"phone" db:"phone"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
}

type UserGet struct {
	Firstname  string `json:"firstname" db:"firstname"`
	Lastname   string `json:"lastname" db:"lastname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Phone      string `json:"phone" db:"phone"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"-" db:"password"`
}

type UserCheck struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}
