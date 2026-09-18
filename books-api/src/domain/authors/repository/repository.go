package repository

import (
	"books-api/domain/authors/models"
	"context"
	"database/sql"
	"errors"
)

type AuthorRepository interface {
	// GetAll returns all authors
	GetAll(ctx context.Context) ([]models.Author, error)
	// GetByID returns a slice of authors with one Author if found, otherwise the slice will be empty
	GetByID(ctx context.Context, ID int32) ([]models.Author, error)
}

type InMemoryRepository struct{}

func (repo InMemoryRepository) GetAll(ctx context.Context) ([]models.Author, error) {
	return models.AUTHORS, nil
}

func (repo InMemoryRepository) GetByID(ctx context.Context, ID int32) ([]models.Author, error) {
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

type PostgresRepository struct {
	db *sql.DB
}

func (repo PostgresRepository) GetAll(ctx context.Context) ([]models.Author, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresRepository) GetByID(ctx context.Context, ID int32) ([]models.Author, error) {
	return nil, errors.New("not implemented")
}

func NewPostgresRepository(db *sql.DB) AuthorRepository {
	return &PostgresRepository{db}
}
