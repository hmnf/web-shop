package main

import (
	"log"

	"github.com/hmnf/web_shop/services"

	"github.com/Moranilt/rou"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const JWTSecret = "jjqqwewewerereresadf"

type Repository struct {
	dbConnection *sqlx.DB
	services     *ServicesList
}

type ServicesList struct {
	Authorization services.AuthMethods
	User          services.UserMethods
	Cards         services.CardsMethods
}

func (r *Repository) Login(ctx *rou.Context) {

}

func (r *Repository) UserInfo(ctx *rou.Context) {

}

func main() {
	router := rou.NewRouter()
	conn, err := sqlx.Connect("postgres", "user=root password=123456 dbname=web_shop sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	cards := services.NewCardsService(conn)
	user := services.NewUserService(conn)
	authorization := services.NewAuthorizationService(conn, JWTSecret, user)
	middleware := services.NewMiddlewareService(conn, authorization)

	repository := Repository{
		dbConnection: conn,
		services: &ServicesList{
			Authorization: authorization,
			User:          user,
			Cards:         cards,
		},
	}

	router.Post("/login", repository.Login)
	router.Get("/user", repository.UserInfo)

	log.Fatal(router.RunServer(":8080"))
}
