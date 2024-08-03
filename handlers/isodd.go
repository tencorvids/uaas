package handlers

import (
	"net/http"
	"strconv"
	"uaas/utilities"

	"github.com/labstack/echo/v4"
)

func IsOddHandler(c echo.Context) error {
	input, _ := strconv.Atoi(c.QueryParam("number"))

	result := utilities.IsOdd(input)
	return c.JSON(http.StatusOK, map[string]bool{"result": result})
}
