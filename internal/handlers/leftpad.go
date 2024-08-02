package handlers

import (
	"net/http"
	"strconv"
	"uaas/internal/utils"

	"github.com/labstack/echo/v4"
)

func LeftPadHandler(c echo.Context) error {
	input := c.QueryParam("string")
	if input == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "string query parameter is required"})
	}

	padding, _ := strconv.Atoi(c.QueryParam("padding"))
	padChar := c.QueryParam("padchar")

	result := utils.LeftPad(input, padding, padChar)
	return c.JSON(http.StatusOK, map[string]string{"result": result})
}
