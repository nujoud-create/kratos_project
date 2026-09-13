package main

import (
	"log"
	"net/http"

	"KratosNew/internal/biz"
	"KratosNew/internal/data"
	"KratosNew/internal/server"
)

func main() {
	repo, err := data.NewUserRepo()
	if err != nil {
		log.Fatal(err)
	}

	uc := biz.NewUserUsecase(repo)
	api := server.NewHTTPServer(uc)

	http.HandleFunc("/users", api.Users)
	http.HandleFunc("/users/", api.Users)

	log.Println("http://localhost:8000/users")

	log.Fatal(http.ListenAndServe(":8000", nil))
}