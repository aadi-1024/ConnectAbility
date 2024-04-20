package handlers

import (
	"github.com/aadi-1024/ConnectAbility/models"
	"github.com/aadi-1024/ConnectAbility/pkg/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func NewChat(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)

		uid2, _ := strconv.Atoi(c.FormValue("uid"))

		chat := models.Chat{
			Member1: uid,
			Member2: uid2,
		}

		if err := db.CreateChat(chat); err != nil {
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

func SendMessage(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)
		chatId, _ := strconv.Atoi(c.FormValue("chat"))
		content := c.FormValue("content")

		msgId, err := db.SendMessage(chatId, uid, content)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: msgId,
		})
	}
}

func GetMessages(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("id").(int)

		chatId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid id",
				Content: err.Error(),
			})
		}

		data, err := db.GetMessages(chatId, uid, 50)
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
