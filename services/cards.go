package services

import "github.com/jmoiron/sqlx"

type Cards struct {
	db *sqlx.DB
}

type Card struct {
	Number    string
	Mask      string
	CVC       int
	Balance   float64
	StateName string
}

type CardsMethods interface {
	GetCards()
	Create()
	Delete()
}

func NewCardsService(db *sqlx.DB) CardsMethods {
	return &Cards{db: db}
}

func (cards *Cards) GetCards() {

}

func (cards *Cards) Create() {

}

func (cards *Cards) Delete() {

}
