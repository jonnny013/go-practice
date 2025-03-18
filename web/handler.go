package web

import (
	"github.com/go-chi/chi"
	goshoppingstore "github.com/jonnny013/go-practice"
)

type Handler struct {
	*chi.Mux
	store goshoppingstore.Store
}