package adapters

import (
	"context"
	"io/ioutil"

	"github.com/go-openapi/runtime/middleware"

	"github.com/joanassum/my-bookstore/bookstore-books/domain"
	"github.com/joanassum/my-bookstore/bookstore-books/models"
	"github.com/joanassum/my-bookstore/bookstore-books/restapi/operations/book_image"
	"github.com/joanassum/my-bookstore/bookstore-books/restapi/operations/books"
)

// BooksHandler handles http requests related to books.
type BooksHandler interface {
	GetAllBooks(params books.GetBooksParams) middleware.Responder
	GetBookImage(params book_image.GetBookIDImageParams) middleware.Responder
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
		internalServerError := "Internal Server Error"
		return books.NewGetBooksDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &internalServerError,
		})
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

func (h *booksHandler) GetBookImage(params book_image.GetBookIDImageParams) middleware.Responder {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", params.ID)
	img, err := h.Ucase.GetBookImage(ctx)
	if err != nil {
		if err.Error() == "not found" || err.Error() == "invalid id" {
			notFoundError := "ID not found"
			return book_image.NewGetBookIDImageNotFound().WithPayload(&models.Error{
				Code:    404,
				Message: &notFoundError,
			})
		}
		internalServerError := "Internal Server Error"
		return book_image.NewGetBookIDImageDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &internalServerError,
		})
	}
	return book_image.NewGetBookIDImageOK().WithPayload(ioutil.NopCloser(img))
}
