package main

import (
	"github.com/aadi-1024/ConnectAbility/pkg/handlers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func addRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.POST("/register", handlers.RegisterUserHandler(app.Db))
	e.POST("/login", handlers.LoginUserHandler(app.Db, app.JwtSecret, app.JwtExpiry))
	e.Static("/static", "static/")
}
