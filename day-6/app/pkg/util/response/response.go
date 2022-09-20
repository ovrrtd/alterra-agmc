package response

import "github.com/labstack/echo/v4"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithJSON(c echo.Context, message string, data interface{}, code int) error {
	return c.JSON(code, Response{
		Message: message,
		Data:    data,
	})
}
