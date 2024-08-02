package middleware

import (
	"strings"
	"uaas/counter"

	"github.com/labstack/echo/v4"
)

func CounterMiddleware(c *counter.Counter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			path := ctx.Path()
			parts := strings.Split(path, "/")
			if len(parts) > 1 {
				endpointName := parts[1]
				c.Increment(endpointName)
			}

			return next(ctx)
		}
	}
}
