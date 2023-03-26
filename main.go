package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TestUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// response, err := http.Get("https://httpbin.org/get")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bytes, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(bytes))

	user := TestUser{Name: "Joe", Email: "joe@gmail.com"}

	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(userBytes))

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
