package repository

import (
	"context"

	"books-api/authors/models"
)

type AuthorRepository interface {
	// GetAll returns all authors
	GetAll(ctx context.Context) ([]models.Author, error)
	// GetByID returns a slice of authors with one Author if found, otherwise the slice will be empty
	GetByID(ctx context.Context, ID string) ([]models.Author, error)
}

type InMemoryRepository struct{}

func (repo InMemoryRepository) GetAll(ctx context.Context) ([]models.Author, error) {
	return models.AUTHORS, nil
}

func (repo InMemoryRepository) GetByID(ctx context.Context, ID string) ([]models.Author, error) {
	for _, user := range models.AUTHORS {
		if user.ID == ID {
			return []models.Author{user}, nil
		}
	}
	return []models.Author{}, nil
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}
