// Package authors implements the Authors resource
package authors

import (
	"books-api/authors/handlers"
	"books-api/authors/repository"

	"github.com/go-chi/chi/v5"
)

func Resource() chi.Router {
	repo := repository.NewInMemoryRepository()
	r := chi.NewRouter()
	r.Get("/", handlers.GetAuthorsHandler(repo))
	r.Get("/{id}", handlers.GetAuthorByIDHandler(repo))

	return r
}
