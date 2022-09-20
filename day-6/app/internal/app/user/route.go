package user

import (
	"agmc/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.CreateUser)
	g.POST("/login", h.Login)

	g.GET("", middleware.UserJwt(h.GetUsers))
	g.GET("/:id", middleware.UserJwt(h.GetUserByID))
	g.PUT("/:id", middleware.UserJwt(h.UpdateUserByID))
	g.DELETE("/:id", middleware.UserJwt(h.DeleteUserByID))
}
