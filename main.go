package main

import (
	"log"

	"github.com/hmnf/web_shop/services"

	"github.com/Moranilt/rou"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	dbConnection *sqlx.DB
	services     *ServicesList
}

type ServicesList struct {
	Authorization services.AuthMethods
}

func main() {
	router := rou.NewRouter()
	conn, err := sqlx.Connect("postgres", "user=root password=123456 dbname=web_shop sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
