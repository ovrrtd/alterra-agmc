package book

import (
	"agmc/internal/dto"
	"agmc/internal/factory"
	"agmc/pkg/util/response"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) GetBooks(c echo.Context) error {
	books, code, err := h.service.GetBooks(c.Request().Context())

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success get all books", books, code)
}

func (h *handler) GetBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	book, code, err := h.service.GetBookByID(c.Request().Context(), int64(id))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success get book by id", book, code)
}

func (h *handler) CreateBook(c echo.Context) error {
	b := new(dto.CreateBookRequest)
	if err := c.Bind(b); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(b); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	book, code, err := h.service.CreateBook(c.Request().Context(), *b)

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success create book", book, code)
}

func (h *handler) DeleteBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	code, err := h.service.DeleteBookByID(c.Request().Context(), int64(id))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success delete book by id", nil, code)
}

func (h *handler) UpdateBookByID(c echo.Context) error {
	b := new(dto.UpdateBookRequest)

	if err := c.Bind(b); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(b); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	book, code, err := h.service.UpdateBookByID(c.Request().Context(), int64(id), *b)

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success update book by id", book, code)
}
