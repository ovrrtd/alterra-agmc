package route

import (
	"agmc/controller"
	"agmc/util"

	m "agmc/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	v1 := e.Group("/v1")
	{

		v1BookRoutes := v1.Group("/books")
		// without auth
		v1BookRoutes.GET("", controller.GetBooks)
		v1BookRoutes.GET("/:id", controller.GetBookByID)
		// with auth
		v1BookRoutes.POST("", m.UserJwt(controller.CreateBook))
		v1BookRoutes.PUT("/:id", m.UserJwt(controller.UpdateBookByID))
		v1BookRoutes.DELETE("/:id", m.UserJwt(controller.DeleteBookByID))
	}

	{

		v1UserRoutes := v1.Group("/users")
		v1UserRoutes.POST("", controller.CreateUser)
		v1UserRoutes.POST("/login", controller.Login)

		v1UserRoutes.GET("", m.UserJwt(controller.GetUsers))
		v1UserRoutes.GET("/:id", m.UserJwt(controller.GetUserByID))
		v1UserRoutes.PUT("/:id", m.UserJwt(controller.UpdateUserByID))
		v1UserRoutes.DELETE("/:id", m.UserJwt(controller.DeleteUserByID))
	}

	return e
}
