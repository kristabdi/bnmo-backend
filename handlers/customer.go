package handlers

import (
	"net/http"
	"strconv"

	"github.com/kristabdi/bnmo-backend/controllers"
	"github.com/kristabdi/bnmo-backend/middleware"
	"github.com/labstack/echo/v4"
)

func GetInfo(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	user, err := controllers.UserGetByUsername(cc.Username)
	if err != nil {
		return cc.JSON(http.StatusInternalServerError, err)
	}

	user.NoSensitive()
	return cc.JSON(http.StatusOK, user)
}

func GetHistory(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	historyType := cc.Param("historyType")
	page, err := strconv.ParseInt(cc.QueryParam("page"), 10, 64)
	if err != nil {
		return cc.String(http.StatusBadRequest, "Invalid page")
	}

	page_size, err := strconv.ParseInt(c.QueryParam("page_size"), 10, 64)
	if err != nil {
		return cc.String(http.StatusBadRequest, "Invalid page size")
	}

	if historyType == "request" {
		requests, err := controllers.RequestGetBatch(page, page_size, cc.ID)
		if err != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		if len(requests) == 0 {
			return cc.NoContent(http.StatusNoContent)
		}
		return cc.JSON(http.StatusOK, requests)
	} else if historyType == "transaction" {
		requests, err := controllers.TransactionGetBatch(page, page_size, cc.ID)
		if err != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		if len(requests) == 0 {
			return cc.NoContent(http.StatusNoContent)
		}
		return cc.JSON(http.StatusOK, requests)
	} else {
		return cc.NoContent(http.StatusBadRequest)
	}
}
