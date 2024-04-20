package handlers

import (
	"github.com/aadi-1024/ConnectAbility/models"
	"github.com/aadi-1024/ConnectAbility/pkg/database"
	"github.com/aadi-1024/ConnectAbility/pkg/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func RegisterUserHandler(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		//resume
		r, err := c.FormFile("resume")
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "resume",
				Content: err.Error(),
			})
		}
		resumeLink, err := util.CreateFile(r)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "resume",
				Content: err.Error(),
			})
		}

		//pfp
		p, err := c.FormFile("profile_pic")
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "pfp",
				Content: err.Error(),
			})
		}
		pfpLink, err := util.CreateFile(p)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "pfp",
				Content: err.Error(),
			})
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), -1)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "password error",
				Content: err.Error(),
			})
		}

		user := models.User{
			Email:           c.FormValue("email"),
			Password:        string(hash),
			PhoneNo:         c.FormValue("phone_no"),
			FirstName:       c.FormValue("first_name"),
			LastName:        c.FormValue("last_name"),
			About:           c.FormValue("about"),
			ProfilePic:      pfpLink,
			ResumeLink:      resumeLink,
			Github:          c.FormValue("github"),
			Linkedin:        c.FormValue("linkedin"),
			Website:         c.FormValue("website"),
			LocationArea:    c.FormValue("location_area"),
			LocationCity:    c.FormValue("location_city"),
			LocationCountry: c.FormValue("location_country"),
			LocationPin:     c.FormValue("location_pin"),
		}

		err = db.RegisterUser(&user)
		if err != nil {
			_ = os.Remove("static/" + pfpLink)
			_ = os.Remove("static/" + resumeLink)
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "db error",
				Content: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
		})
	}
}

func LoginUserHandler(db *database.Database, secret []byte, exp time.Duration) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		usr := models.User{Email: email, Password: password}
		id, err := db.LoginUser(&usr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid credentials",
				Content: err.Error(),
			})
		}

		token, err := util.GenerateJwtToken(id, secret, exp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "something went wrong",
				Content: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: token,
		})
	}
}
