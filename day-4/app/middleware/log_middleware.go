package middleware

import (
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(m.LoggerWithConfig(m.LoggerConfig{
		Format: `[${time_rfc3339}]method=${method}, uri=${uri}, status=${status} ${latency_human}` + "\n",
	}))
}
