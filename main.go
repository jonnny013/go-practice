package main

import (
	"net/http"

	"github.com/jonnny013/go-practice/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
