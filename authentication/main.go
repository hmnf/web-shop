package authentication

import (
	"log"
	"net/http"

	"github.com/Moranilt/rou"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	dbConnection   *sqlx.DB
	authentication AuthMethods
}

func (r *Repository) Login(ctx *rou.Context) {
	user, err := r.authentication.Login(ctx.ResponseWriter(), "joe@mail.com", "123456")

	if err != nil {
		ctx.ErrorJSONResponse(http.StatusBadRequest, err.Error())
		return
	}

	ctx.SuccessJSONResponse(user)
}

func main() {
	router := rou.NewRouter()
	conn, err := sqlx.Connect("postgres", "user=root password=123456 dbname=transactions sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	JWTSecret := "jsahdjabns121='31/"
	user := NewUserService(conn)
	authentication := NewAuthenticationService(conn, JWTSecret, user)
	// middleware := NewMiddlewareService(conn, authentication)

	repository := Repository{
		dbConnection:   conn,
		authentication: authentication,
	}

	router.Post("/login", repository.Login)

	log.Fatal(router.RunServer(":8080"))
}
