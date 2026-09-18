package server

import (
	"books-api/domain/authors"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// Global Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// API Routes
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/authors", authors.Resource(db))
	})

	return r
}
