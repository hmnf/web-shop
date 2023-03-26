package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sqlx.DB
}


func (repo *Repository) CreateUser(w http.ResponseWriter, r *http.Request) {
	result, err := decodeBody[UserCreate](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(result.Password), 14)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result.Password = string(hash)

	_, err = repo.db.NamedExec(`INSERT INTO users (firstname, lastname, patronymic, phone, email, password, role) 
	VALUES (:firstname, :lastname, :patronymic, :phone, :email, :password, '3')`, result)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func (repo *Repository) GetUser(w http.ResponseWriter, r *http.Request) {
	result, err := decodeBody[UserCheck](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var user UserGet

	err = repo.db.Get(&user, "SELECT firstname, lastname, patronymic, phone, email, password FROM users WHERE email=$1", result.Email)
	if err == sql.ErrNoRows {
		log.Println("user not found")
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(result.Password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(user)
}

func main() {
	conn, err := sqlx.Connect("postgres", "dbname=users host=localhost port=5432 user=root password=123456 sslmode=disable")

	if err != nil {
		log.Fatal("Counldn't connect to db users: ", err)
	}

	defer conn.Close()

	repo := &Repository{
		db: conn,
	}

	router := mux.NewRouter()
	router.HandleFunc("/create", repo.CreateUser).Methods("POST")
	router.HandleFunc("/get", repo.GetUser).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", router))
}
