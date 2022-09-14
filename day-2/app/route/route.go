package route

import (
	"day-two/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	{
		v1BookRoutes := v1.Group("/books")

		v1BookRoutes.GET("", controller.GetBooks)
		v1BookRoutes.POST("", controller.CreateBook)
		v1BookRoutes.GET("/:id", controller.GetBookByID)
		v1BookRoutes.PUT("/:id", controller.UpdateBookByID)
		v1BookRoutes.DELETE("/:id", controller.DeleteBookByID)
	}

	{
		v1UserRoutes := v1.Group("/users")

		v1UserRoutes.GET("", controller.GetUsers)
		v1UserRoutes.POST("", controller.CreateUser)
		v1UserRoutes.GET("/:id", controller.GetUserByID)
		v1UserRoutes.PUT("/:id", controller.UpdateUserByID)
		v1UserRoutes.DELETE("/:id", controller.DeleteUserByID)
	}

	return e
}
