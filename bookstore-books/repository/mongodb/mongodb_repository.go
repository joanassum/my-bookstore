package mongodb

import (
	"bytes"
	"context"
	"errors"
	"io"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/joanassum/my-bookstore/bookstore-books/domain"
)

const DB_NAME = "my-bookstore"
const BOOKS_COLLECTION = "books"

type mongoDBRepository struct {
	DB *mgo.Session
}

// NewMongoDBRepository returns a mongoDBRepository.
func NewMongoDBRepository(db *mgo.Session) domain.BooksRepository {
	return &mongoDBRepository{DB: db}
}

type book struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `json:"title"`
	Author      string        `json:"author"`
	Description string        `json:"description"`
	Price       float64       `json:"price"`
	Image       image         `json:"image"`
}

type image struct {
	Data bson.Binary `json:"data"`
}

// GetAllBooks returns a list of books from the database.
func (m *mongoDBRepository) GetAllBooks(ctx context.Context) ([]*domain.Book, error) {
	books := make([]*domain.Book, 0)
	iter := m.DB.DB(DB_NAME).C(BOOKS_COLLECTION).Find(nil).Iter()
	var b book
	for iter.Next(&b) {
		books = append(books, &domain.Book{
			ID:     string(b.ID),
			Title:  b.Title,
			Author: b.Author,
			Price:  b.Price,
		})
	}
	return books, nil
}

// GetBookImage returns the byte representation of image of the book
func (m *mongoDBRepository) GetBookImage(ctx context.Context) (io.Reader, error) {
	id := ctx.Value("id")
	idStr, ok := id.(string)
	if !ok {
		return nil, errors.New("cannot convert ID to string")
	}
	if !bson.IsObjectIdHex(idStr) {
		return nil, errors.New("invalid ID")
	}
	var result book
	err := m.DB.DB(DB_NAME).C(BOOKS_COLLECTION).FindId(bson.ObjectIdHex(idStr)).One(&result)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(result.Image.Data.Data), nil
}
