package book

import (
	"agmc/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {

	// without auth
	g.GET("", h.GetBooks)
	g.GET("/:id", h.GetBookByID)
	// with auth
	g.POST("", middleware.UserJwt(h.CreateBook))
	g.PUT("/:id", middleware.UserJwt(h.UpdateBookByID))
	g.DELETE("/:id", middleware.UserJwt(h.DeleteBookByID))
}
