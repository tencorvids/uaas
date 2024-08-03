package main

import (
	"log/slog"
	"os"
	"time"
	"uaas/config"
	"uaas/counter"
	"uaas/handlers"
	m "uaas/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	countsFile = "/app/data/counts.json"
)

func main() {
	cfg := config.Load()

	c := counter.NewCounter(countsFile, 5*time.Minute)

	e := echo.New()
	e.Use(m.RequestLogger(cfg.Logger))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(m.CounterMiddleware(c))

	handlers.RegisterRoutes(e)

	e.GET("/counts", func(ctx echo.Context) error {
		return ctx.JSON(200, c.GetAll())
	})

	slog.Info("Starting server", "port", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
