package main

import (
	"log/slog"
	"os"
	"uaas/config"
	"uaas/handlers"
	m "uaas/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.Load()

	e := echo.New()
	e.Use(m.RequestLogger(cfg.Logger))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handlers.RegisterRoutes(e)

	slog.Info("Starting server", "port", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
