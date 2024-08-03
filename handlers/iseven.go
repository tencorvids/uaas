package handlers

import (
	"net/http"
	"strconv"
	"uaas/utilities"

	"github.com/labstack/echo/v4"
)

func IsEvenHandler(c echo.Context) error {
	input, _ := strconv.Atoi(c.QueryParam("number"))

	result := utilities.IsEven(input)
	return c.JSON(http.StatusOK, map[string]bool{"result": result})
}
