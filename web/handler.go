package web

import (
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	goshoppingstore "github.com/jonnny013/go-practice"
)

func NewHandler(store goshoppingstore.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	h.Route("/items", func(r chi.Router) {
		r.Get("/", h.ItemsList())
	})

	return h
}

type Handler struct {
	*chi.Mux
	store goshoppingstore.Store
}

const threadsListHTML = `
<h1>Items</h1>
<dl>
{{range .Items}}
	<dt>{{.Name}}</dt>
	<dd>{{.Description}}</dd>
{{end}}
</dl>
`

func (h *Handler) ItemsList() http.HandlerFunc {
	type data struct {
		Items []goshoppingstore.Item
	}
	tmpl := template.Must(template.New("th").Parse(threadsListHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Items()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data{Items: tt})
	}
}
