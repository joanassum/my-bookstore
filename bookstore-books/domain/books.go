package domain

import (
	"context"
	"io"
)

//go:generate mockgen -destination=./mocks/mock_books.go -package=mocks -source=./books.go BooksRepository,BooksUsecase

type Book struct {
	ID     string
	Title  string
	Author string
	Price  float64
}

type BooksRepository interface {
	GetAllBooks(ctx context.Context) ([]*Book, error)
	GetBookImage(ctx context.Context) (io.Reader, error)
}

type BooksUsecase interface {
	GetAllBooks(ctx context.Context) ([]*Book, error)
	GetBookImage(ctx context.Context) (io.Reader, error)
}
