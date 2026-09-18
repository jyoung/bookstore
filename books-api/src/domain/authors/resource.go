// Package authors implements the Authors resource
package authors

import (
	handlers2 "books-api/domain/authors/handlers"
	"books-api/domain/authors/repository"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Resource(db *sql.DB) chi.Router {
	repo := repository.NewPostgresRepository(db)
	r := chi.NewRouter()
	r.Get("/", handlers2.GetAuthorsHandler(repo))
	r.Get("/{id}", handlers2.GetAuthorByIDHandler(repo))

	return r
}
