package main

import (
	"log"
	"net/http"

	"github.com/jonnny013/go-practice/postgres"
	"github.com/jonnny013/go-practice/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
