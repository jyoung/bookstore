package handlers

import (
	"books-api/respond"
	"net/http"

	"books-api/authors/repository"
)

func GetAuthorsHandler(repo repository.AuthorRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := repo.GetAll(r.Context())
		if err != nil {
			respond.Error(w, r, respond.Unhandled("failed to fetch authors"))
			return
		}

		respond.OK(w, authors)
	}
}
