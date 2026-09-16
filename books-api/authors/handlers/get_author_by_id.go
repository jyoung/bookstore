package handlers

import (
	"books-api/respond"
	"fmt"
	"net/http"

	"books-api/authors/repository"

	"github.com/go-chi/chi/v5"
)

func GetAuthorByIDHandler(repo repository.AuthorRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorID := chi.URLParam(r, "id")
		author, err := repo.GetByID(r.Context(), authorID)
		if err != nil {
			respond.Error(w, r, respond.Unhandled("failed to fetch author"))
			return
		}

		if len(author) == 0 {
			respond.Error(w, r, respond.NotFound(fmt.Sprintf("author %s", authorID)))
			return
		}

		respond.OK(w, author)
	}
}
