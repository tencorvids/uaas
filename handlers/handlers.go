package handlers

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/left-pad", LeftPadHandler)
	e.GET("/right-pad", RightPadHandler)
	e.GET("/is-even", IsEvenHandler)
	e.GET("/is-odd", IsOddHandler)
}
