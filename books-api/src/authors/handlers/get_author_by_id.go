package handlers

import (
	"books-api/respond"
	"fmt"
	"net/http"
	"strconv"

	"books-api/authors/repository"

	"github.com/go-chi/chi/v5"
)

func GetAuthorByIDHandler(repo repository.AuthorRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorIDParam := chi.URLParam(r, "id")
		authorID, err := strconv.ParseInt(authorIDParam, 10, 32)
		if err != nil {
			respond.Error(w, r, respond.BadRequest("invalid author id"))
			return
		}

		author, err := repo.GetByID(r.Context(), int32(authorID))
		if err != nil {
			respond.Error(w, r, respond.Unhandled("failed to fetch author"))
			return
		}

		if len(author) == 0 {
			respond.Error(w, r, respond.NotFound(fmt.Sprintf("author %s", authorIDParam)))
			return
		}

		respond.Ok(w, r, respond.Single(author))
	}
}
