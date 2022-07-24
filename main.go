package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kristabdi/bnmo-backend/handlers"
	"github.com/kristabdi/bnmo-backend/utils"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalln("Env load failed")
	}

	if err = utils.Db.InitDB(); err != nil {
		log.Fatalln("DB Connection error")
	}

	if err = utils.Db.InitSeeding(); err != nil {
		log.Fatalln("Seeding error")
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	auth := e.Group("/auth")
	auth.POST("/register", handlers.Registration)
	auth.POST("/login", handlers.Login)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
