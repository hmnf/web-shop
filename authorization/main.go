package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func (repo *Repository) Login(w http.ResponseWriter, r *http.Request) {
	userLogin, err := decodeBody[UserLogin](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

}

func main() {
	router := mux.NewRouter()
	conn, err := sqlx.Connect("postgres", "user=root password=123456 dbname=authorization sslmode=disable")

	if err != nil {
		log.Fatal("Counldn't connect to db users: ", err)
	}

	defer conn.Close()

	repo := &Repository{
		db: conn,
	}

	router.HandleFunc("/login", repo.Login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))
}
