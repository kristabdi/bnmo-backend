package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kristabdi/bnmo-backend/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalln("Env load failed")
	}

	err = utils.InitDB()
	if err != nil {
		log.Fatalln("DB Connection error")
	}

	err = utils.InitSeeding()
	if err != nil {
		log.Fatalln("Seeding error")
	}

	fmt.Println("Db Connected")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
