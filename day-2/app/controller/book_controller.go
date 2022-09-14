package controller

import (
	"day-two/lib/static"
	"day-two/model"
	"day-two/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetBooks(c echo.Context) error {
	books, code, err := static.GetBooks(c.Request().Context())

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	res := make([]Book, len(books))

	for i, book := range books {
		res[i] = Book{
			ID:   book.ID,
			Name: book.Name,
		}
	}

	return util.ResponseWithJSON(c, "success get all books", res, code)
}

func GetBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	book, code, err := static.GetBookByID(c.Request().Context(), int64(id))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success get book by id", Book{
		ID:   book.ID,
		Name: book.Name,
	}, code)
}

type CreateBookRequest struct {
	Name string `json:"name" binding:"required"`
}

func CreateBook(c echo.Context) error {
	b := new(CreateBookRequest)
	if err := c.Bind(b); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	book, code, err := static.CreateBook(c.Request().Context(), model.Book{
		Name: b.Name,
	})

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success create book", Book{
		ID:   book.ID,
		Name: book.Name,
	}, code)
}

func DeleteBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	code, err := static.DeleteBookByID(c.Request().Context(), int64(id))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success delete book by id", nil, code)
}

type UpdateBookRequest struct {
	Name string `json:"name" binding:"required"`
}

func UpdateBookByID(c echo.Context) error {
	b := new(CreateBookRequest)

	if err := c.Bind(b); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	book, code, err := static.UpdateBookByID(c.Request().Context(), int64(id), model.Book{
		Name: b.Name,
	})

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success update book by id", Book{
		ID:   book.ID,
		Name: book.Name,
	}, code)
}
