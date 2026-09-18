package handlers

import (
	"books-api/domain/authors/repository"
	"books-api/respond"
	"net/http"
)

func GetAuthorsHandler(repo repository.AuthorRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := repo.GetAll(r.Context())
		if err != nil {
			respond.Error(w, r, respond.Unhandled("failed to fetch authors"))
			return
		}

		respond.Ok(w, r, respond.List(authors, 1, 2, 1))
	}
}
