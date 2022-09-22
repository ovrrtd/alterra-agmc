package repository

import (
	"agmc/internal/model"
	"context"
	"errors"
	"net/http"
)

type Book interface {
	GetBooks(ctx context.Context) ([]model.Book, int, error)
	GetBookByID(ctx context.Context, ID int64) (*model.Book, int, error)
	CreateBook(ctx context.Context, b model.Book) (model.Book, int, error)
	DeleteBookByID(ctx context.Context, ID int64) (int, error)
	UpdateBookByID(ctx context.Context, ID int64, data model.Book) (*model.Book, int, error)
}

type book struct {
	Books []model.Book
}

func NewBook() Book {
	books := []model.Book{
		{
			ID:   1,
			Name: "Book 1",
		},
		{
			ID:   2,
			Name: "Book 2",
		},
	}
	return &book{
		Books: books,
	}
}

func (b *book) GetBooks(ctx context.Context) ([]model.Book, int, error) {
	return b.Books, http.StatusOK, nil
}

func (b *book) GetBookByID(ctx context.Context, ID int64) (*model.Book, int, error) {
	for _, book := range b.Books {
		if book.ID == ID {
			return &book, http.StatusOK, nil
		}
	}

	return nil, http.StatusNotFound, errors.New("book not found")
}

func (b *book) CreateBook(ctx context.Context, book model.Book) (model.Book, int, error) {
	id := b.Books[len(b.Books)-1].ID + 1
	book.ID = id

	b.Books = append(b.Books, book)

	return book, http.StatusCreated, nil
}

func (b *book) DeleteBookByID(ctx context.Context, ID int64) (int, error) {
	var temp []model.Book
	found := false
	for _, book := range b.Books {
		if book.ID == ID {
			found = true
			continue
		}
		temp = append(temp, book)
	}

	if !found {
		return http.StatusNotFound, errors.New("book not found")
	}

	b.Books = temp

	return http.StatusOK, nil
}

func (b *book) UpdateBookByID(ctx context.Context, ID int64, data model.Book) (*model.Book, int, error) {
	index := -1
	for i, book := range b.Books {
		if book.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, http.StatusNotFound, errors.New("book not found")

	}

	b.Books[index].Name = data.Name

	return &b.Books[index], http.StatusOK, nil
}
