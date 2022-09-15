package static

import (
	"agmc/model"
	"context"
	"errors"
	"net/http"
)

var books []model.Book = []model.Book{
	{
		ID:   1,
		Name: "Book 1",
	},
	{
		ID:   2,
		Name: "Book 2",
	},
}

func GetBooks(ctx context.Context) ([]model.Book, int, error) {
	return books, http.StatusOK, nil
}

func GetBookByID(ctx context.Context, ID int64) (*model.Book, int, error) {
	for _, book := range books {
		if book.ID == ID {
			return &book, http.StatusOK, nil
		}
	}

	return nil, http.StatusNotFound, errors.New("book not found")
}

func CreateBook(ctx context.Context, b model.Book) (model.Book, int, error) {
	id := books[len(books)-1].ID + 1
	b.ID = id

	books = append(books, b)

	return b, http.StatusCreated, nil
}

func DeleteBookByID(ctx context.Context, ID int64) (int, error) {
	var temp []model.Book
	found := false
	for _, book := range books {
		if book.ID == ID {
			found = true
			continue
		}
		temp = append(temp, book)
	}

	if !found {
		return http.StatusNotFound, errors.New("book not found")
	}

	books = temp

	return http.StatusOK, nil
}

func UpdateBookByID(ctx context.Context, ID int64, data model.Book) (*model.Book, int, error) {
	index := -1
	for i, book := range books {
		if book.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, http.StatusNotFound, errors.New("book not found")

	}

	books[index].Name = data.Name

	return &books[index], http.StatusOK, nil
}
