package domain

import "context"

type Book struct {
	ID     string
	Title  string
	Author string
	Price  float64
}

type BooksRepository interface {
	GetAllBooks(ctx context.Context) ([]*Book, error)
}

type BooksUsecase interface {
	GetAllBooks(ctx context.Context) ([]*Book, error)
}
