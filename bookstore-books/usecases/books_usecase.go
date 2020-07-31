package usecases

import (
	"context"

	"github.com/joanassum/my-bookstore/bookstore-books/domain"
)

type bookUsecase struct {
	Repo domain.BooksRepository
}

// NewBookUsecase returns a new bookUsecase.
func NewBookUsecase(repo domain.BooksRepository) domain.BooksUsecase {
	return &bookUsecase{Repo: repo}
}

// GetAllBooks returns a list of books
func (b *bookUsecase) GetAllBooks(ctx context.Context) ([]*domain.Book, error) {
	return b.Repo.GetAllBooks(ctx)
}
