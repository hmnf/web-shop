package main

import "github.com/jmoiron/sqlx"

type Middleware struct {
	db   *sqlx.DB
	auth AuthMethods
}

func NewMiddlewareService(db *sqlx.DB, auth AuthMethods) *Middleware {
	return &Middleware{
		db:   db,
		auth: auth,
	}
}
