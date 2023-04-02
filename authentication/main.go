package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var redisContext = context.Background()

func CreateTokens(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:8002",
		Password: "",
		DB:       0,
	})

	auth := &Authentication{
		Secret:       secret,
		RedisClient:  rdb,
		RedisContext: redisContext,
	}

	user, err := decodeBody[UserDetails](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = auth.CreateTokens(user.Id, user.Role)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// json.NewEncoder(w).Encode(td.AccessToken)
}

func UpdateTokens(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/create", CreateTokens).Methods("POST")
	router.HandleFunc("/update", UpdateTokens).Methods("GET")

	log.Fatal(http.ListenAndServe(":8002", router))
}
