package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Auth *Authentication
}

var redisContext = context.Background()

func (repo *Repository) CreateTokens(w http.ResponseWriter, r *http.Request) {
	user, err := decodeBody[UserDetails](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	tokens, err := repo.Auth.CreateTokens(user.Id, user.Role)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(tokens)
}

func (repo *Repository) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	tokens, err := decodeBody[Tokens](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	refreshedTokens, err := repo.Auth.RefreshTokens(tokens.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(refreshedTokens)
}

func main() {
	router := mux.NewRouter()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	auth := &Authentication{
		Secret:       secret,
		RedisClient:  rdb,
		RedisContext: redisContext,
	}

	repo := &Repository{
		Auth: auth,
	}

	router.HandleFunc("/create", repo.CreateTokens).Methods("POST")
	router.HandleFunc("/refresh", repo.RefreshTokens).Methods("PUT")
	// router.HandleFunc("/validate", repo.ValidateTokens).Methods("POST")

	log.Fatal(http.ListenAndServe(":8002", router))
}
