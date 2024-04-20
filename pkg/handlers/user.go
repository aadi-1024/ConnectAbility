package handlers

import (
	"github.com/aadi-1024/ConnectAbility/models"
	"github.com/aadi-1024/ConnectAbility/pkg/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InviteHandler(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)

		id, err := strconv.Atoi(c.FormValue("uid"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid id",
				Content: err.Error(),
			})
		}
		teamId, err := strconv.Atoi(c.FormValue("tid"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid id",
				Content: err.Error(),
			})
		}
		owner, err := db.GetTeamOwner(teamId)
		if owner != uid {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "not authorized",
				Content: nil,
			})
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "not authorized",
				Content: err.Error(),
			})
		}

		inv := models.TeamInvite{
			TeamId:    teamId,
			InvitedId: id,
		}
		err = db.CreateTeamInvite(inv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: nil,
		})
	}
}

func ResolveHandler(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)

		invId, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid id",
				Content: err.Error(),
			})
		}
		accept, _ := strconv.ParseBool(c.FormValue("accept"))
		err = db.ResolveTeamInvite(uid, invId, accept)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: nil,
		})
	}
}
