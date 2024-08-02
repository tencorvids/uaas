package middleware

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			logger.Info("request",
				slog.String("method", c.Request().Method),
				slog.String("path", c.Path()),
				slog.Int("status", c.Response().Status),
				slog.Duration("duration", time.Since(start)),
			)
			return err
		}
	}
}
