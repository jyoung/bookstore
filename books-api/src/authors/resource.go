// Package authors implements the Authors resource
package authors

import (
	"books-api/authors/handlers"
	"books-api/authors/repository"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Resource(db *sql.DB) chi.Router {
	repo := repository.NewPostgresRepository(db)
	r := chi.NewRouter()
	r.Get("/", handlers.GetAuthorsHandler(repo))
	r.Get("/{id}", handlers.GetAuthorByIDHandler(repo))

	return r
}
