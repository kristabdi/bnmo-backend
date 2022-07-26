package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kristabdi/bnmo-backend/controllers"
	"github.com/kristabdi/bnmo-backend/middleware"
	"github.com/kristabdi/bnmo-backend/models"
	"github.com/kristabdi/bnmo-backend/utils"
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
			log.Println(err)
			return cc.NoContent(http.StatusInternalServerError)
		}
		if len(requests) == 0 {
			return cc.NoContent(http.StatusNoContent)
		}
		return cc.JSON(http.StatusOK, requests)
	} else if historyType == "transaction" {
		requests, err := controllers.TransactionGetBatch(page, page_size, cc.ID)
		if err != nil {
			log.Println(err)
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

func Withdraw(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	req := new(models.Request)
	if err := c.Bind(req); err != nil {
		return cc.NoContent(http.StatusBadRequest)
	}

	var rate float64

	val, err := utils.Db.Client.Get(utils.Db.Context, req.Currency).Result()
	if err != nil {
		//	Exchange rate api
		url := `https://api.apilayer.com/exchangerates_data/convert?to=IDR&from=` + req.Currency + `&amount=` + strconv.Itoa(int(req.Amount))
		client := &http.Client{}
		request, err2 := http.NewRequest("GET", url, nil)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		request.Header.Add("apikey", os.Getenv("API_KEY"))
		res, err2 := client.Do(request)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		resBody, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}

		var conversion models.Converter
		err2 = json.Unmarshal(resBody, &conversion)
		if err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}

		if err2 = utils.Db.Client.Set(utils.Db.Context, req.Currency, conversion.Info.Rate, 20*time.Minute).Err(); err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}

		rate = conversion.Info.Rate
	} else {
		var err2 error
		rate, err2 = strconv.ParseFloat(val, 64)
		if err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}
	}

	req.UserID = cc.ID
	req.IsAdd = false
	req.Amount = uint64(math.Floor(rate * float64(req.Amount)))

	if req.Amount > cc.Balance {
		return cc.NoContent(http.StatusBadRequest)
	}

	if err := controllers.RequestInsertOne(req); err != nil {
		return cc.NoContent(http.StatusInternalServerError)
	}
	return cc.NoContent(http.StatusOK)
}

func Deposit(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	req := new(models.Request)
	if err := c.Bind(req); err != nil {
		return cc.NoContent(http.StatusBadRequest)
	}
	var rate float64

	val, err := utils.Db.Client.Get(utils.Db.Context, req.Currency).Result()
	if err != nil {
		//	Exchange rate api
		url := `https://api.apilayer.com/exchangerates_data/convert?to=IDR&from=` + req.Currency + `&amount=` + strconv.Itoa(int(req.Amount))
		client := &http.Client{}
		request, err2 := http.NewRequest("GET", url, nil)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		request.Header.Add("apikey", os.Getenv("API_KEY"))
		res, err2 := client.Do(request)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		resBody, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}

		var conversion models.Converter
		err2 = json.Unmarshal(resBody, &conversion)
		if err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}

		if err2 = utils.Db.Client.Set(utils.Db.Context, req.Currency, conversion.Info.Rate, 20*time.Minute).Err(); err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}

		rate = conversion.Info.Rate
	} else {
		var err2 error
		rate, err2 = strconv.ParseFloat(val, 64)
		if err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}
	}
	req.UserID = cc.ID
	req.IsAdd = true
	req.Amount = uint64(math.Floor(rate * float64(req.Amount)))

	if err := controllers.RequestInsertOne(req); err != nil {
		return cc.NoContent(http.StatusInternalServerError)
	}
	return cc.NoContent(http.StatusOK)
}

func Transaction(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	transaction := new(models.Transaction)
	if err := c.Bind(transaction); err != nil {
		return cc.NoContent(http.StatusBadRequest)
	}

	var rate float64

	val, err := utils.Db.Client.Get(utils.Db.Context, transaction.CurrencyFrom).Result()
	if err != nil {
		//	Exchange rate api
		url := `https://api.apilayer.com/exchangerates_data/convert?to=IDR&from=` + transaction.CurrencyFrom + `&amount=` + strconv.Itoa(int(transaction.Amount))
		client := &http.Client{}
		request, err2 := http.NewRequest("GET", url, nil)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		request.Header.Add("apikey", os.Getenv("API_KEY"))
		res, err2 := client.Do(request)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		resBody, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}

		var conversion models.Converter
		err2 = json.Unmarshal(resBody, &conversion)
		if err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}

		if err2 = utils.Db.Client.Set(utils.Db.Context, transaction.CurrencyFrom, conversion.Info.Rate, 20*time.Minute).Err(); err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}

		rate = conversion.Info.Rate
	} else {
		var err2 error
		rate, err2 = strconv.ParseFloat(val, 64)
		if err2 != nil {
			return cc.JSON(http.StatusInternalServerError, err2)
		}
	}

	var destination models.User
	destination, err = controllers.UserGetByUsername(transaction.UsernameTo)

	transaction.IdFrom = cc.ID
	transaction.IdTo = destination.ID
	transaction.UsernameTo = destination.Username
	if err != nil {
		return cc.String(http.StatusNotFound, "Destination Not Found")
	}
	transaction.Amount = uint64(math.Floor(rate * float64(transaction.Amount)))
	if transaction.Amount > cc.Balance {
		return cc.NoContent(http.StatusBadRequest)
	}

	if err = controllers.TransactionInsertOne(transaction); err != nil {
		return cc.NoContent(http.StatusInternalServerError)
	}

	if err = controllers.UserUpdateBalance(transaction.IdFrom, cc.Balance-transaction.Amount); err != nil {
		return cc.NoContent(http.StatusInternalServerError)
	}

	if err = controllers.UserUpdateBalance(transaction.IdTo, destination.Balance+transaction.Amount); err != nil {
		return cc.NoContent(http.StatusInternalServerError)
	}

	return cc.NoContent(http.StatusOK)
}
