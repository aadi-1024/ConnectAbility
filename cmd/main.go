package main

import (
	"github.com/aadi-1024/ConnectAbility/pkg/database"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

var app = &Config{}

func main() {
	app.JwtSecret = []byte("secret")
	app.JwtExpiry = 24 * time.Hour
	d, err := database.InitDb("postgres://postgres:password@localhost:5432/connectability")
	if err != nil {
		log.Fatalln(err)
	}
	app.Db = d

	e := echo.New()
	addRoutes(e)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalln(err)
	}
}
