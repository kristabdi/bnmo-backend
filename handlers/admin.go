package handlers

import (
	"net/http"

	"github.com/kristabdi/bnmo-backend/controllers"
	"github.com/labstack/echo/v4"
)

func VerifyUser(c echo.Context) error {
	username := c.QueryParam("username")
	if err := controllers.UserVerify(username); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "OK")
}

func VerifyRequest(c echo.Context) error {
	return c.String(200, "OK")
}

func GetUsers(c echo.Context) error {
	users, err := controllers.UserGetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(users) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, users)
}
