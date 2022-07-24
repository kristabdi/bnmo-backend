package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kristabdi/bnmo-backend/controllers"
	"github.com/kristabdi/bnmo-backend/models"
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
	var id uuid.UUID
	var err error
	id, err = uuid.Parse(c.QueryParam("id"))

	var req models.Request
	req, err = controllers.RequestGetById(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	var user models.User
	user, err = controllers.UserGetById(req.UserID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if req.IsAdd == false && user.Balance < req.Amount {
		return c.NoContent(http.StatusBadRequest)
	}

	if err = controllers.RequestVerify(id); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if err = controllers.UserUpdateBalance(req.UserID, user.Balance+req.Amount); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(200, "OK")
}

func GetUsers(c echo.Context) error {
	users, err := controllers.UserGetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(users) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}

func GetRequest(c echo.Context) error {
	req, err := controllers.RequestGetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(req) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, req)
}
