package main

import (
	"log"

	"github.com/Moranilt/rou"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

const JWTSecret = "12312321IDFOIDUFAP123"

func main() {
	router := rou.NewRouter()
	conn, err := sqlx.Connect("postgres", "user=root password=123456 dbname=transactions sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := &Repository{
		db: conn,
	}

	// router.Post("/login", repo.Login)

	log.Fatal(router.RunServer(":8080"))
}
