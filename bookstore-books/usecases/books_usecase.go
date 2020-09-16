package usecases

import (
	"context"
	"io"

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

// GetBookImage returns byte representation of the image of the book
func (b *bookUsecase) GetBookImage(ctx context.Context) (io.Reader, error) {
	return b.Repo.GetBookImage(ctx)
}
