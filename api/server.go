package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	goshoppingstore "github.com/jonnny013/go-practice"
)

type Server struct {
	*mux.Router
	shoppingItems []goshoppingstore.Item
}

func NewServer() *Server {
	s := &Server{
		Router:        mux.NewRouter(),
		shoppingItems: []goshoppingstore.Item{},
	}
	s.Router.StrictSlash(true)
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/shopping-items", s.listShoppingItems()).Methods("GET")
	s.HandleFunc("/shopping-items", s.createShoppingItem()).Methods("POST")
	s.HandleFunc("/shopping-items/{id}", s.removeShoppingItem()).Methods("DELETE")
}

func (s *Server) createShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i goshoppingstore.Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.Id = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.shoppingItems {
			if item.Id == id {
				s.shoppingItems = append(s.shoppingItems[:i], s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}
