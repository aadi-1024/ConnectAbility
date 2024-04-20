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

	teams := e.Group("/teams", JwtMiddleware)
	teams.POST("/create", handlers.CreateTeamHandler(app.Db))
	teams.GET("/get", handlers.GetTeams(app.Db))
	teams.GET("/get/:id", handlers.GetTeamById(app.Db))

	user := e.Group("/user", JwtMiddleware)
	user.POST("/invite", handlers.InviteHandler(app.Db))
	user.POST("/resolve-invite", handlers.ResolveHandler(app.Db))

	chat := e.Group("/chat", JwtMiddleware)
	chat.POST("/new-chat", handlers.NewChat(app.Db))
	chat.POST("/send-msg", handlers.SendMessage(app.Db))
	chat.GET("/messages/:id", handlers.GetMessages(app.Db))
	e.Static("/static", "static/")
}
