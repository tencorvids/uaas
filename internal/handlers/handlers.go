package handlers

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/leftpad", LeftPadHandler)
}
