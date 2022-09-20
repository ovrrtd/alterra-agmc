package http

import (
	"agmc/internal/app/book"
	"agmc/internal/app/user"
	"agmc/internal/factory"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	v1 := e.Group("/v1")
	user.NewHandler(f).Route(v1.Group("/users"))
	book.NewHandler(f).Route(v1.Group("/books"))
}
