package adapters

import (
	"context"

	"github.com/go-openapi/runtime/middleware"

	"github.com/joanassum/my-bookstore/bookstore-books/domain"
	"github.com/joanassum/my-bookstore/bookstore-books/models"
	"github.com/joanassum/my-bookstore/bookstore-books/restapi/operations/books"
)

// BooksHandler handles http requests related to books.
type BooksHandler interface {
	GetAllBooks(params books.GetBooksParams) middleware.Responder
}

type booksHandler struct {
	Ucase domain.BooksUsecase
}

// NewBooksHandler returns a new BooksHandler
func NewBooksHandler(ucase domain.BooksUsecase) BooksHandler {
	return &booksHandler{Ucase: ucase}
}

// GetAllBooks response to GET /books request with details of a list of books.
func (h *booksHandler) GetAllBooks(params books.GetBooksParams) middleware.Responder {
	ctx := context.Background()
	bks, err := h.Ucase.GetAllBooks(ctx)
	if err != nil {
		return books.NewGetBooksDefault(500)
	}
	res := make([]*models.Book, 0)
	for _, bk := range bks {
		res = append(res, &models.Book{
			Author: &bk.Author,
			ID:     &bk.ID,
			Title:  &bk.Title,
			Price:  &bk.Price,
		})
	}
	return books.NewGetBooksOK().WithPayload(res)
}
