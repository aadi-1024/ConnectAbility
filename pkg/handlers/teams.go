package handlers

import (
	"github.com/aadi-1024/ConnectAbility/models"
	"github.com/aadi-1024/ConnectAbility/pkg/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func CreateTeamHandler(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)
		team := models.Teams{}

		team.OwnerId = uid
		team.Name = c.FormValue("name")
		team.CreationTime = time.Now()

		tid, err := db.CreateTeam(team)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: tid,
		})
	}
}

func GetTeams(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)
		data, err := db.GetTeams(uid, 10)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: data,
		})
	}
}

func GetTeamById(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "incorrect id",
				Content: err.Error(),
			})
		}

		team, err := db.GetTeamById(uid, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: team,
		})
	}
}
