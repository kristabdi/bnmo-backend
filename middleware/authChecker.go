package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/kristabdi/bnmo-backend/controllers"
	"github.com/kristabdi/bnmo-backend/models"
	"github.com/kristabdi/bnmo-backend/utils"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	models.User
}

func CheckUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})

		claims, ok := token.Claims.(*utils.CustomClaims)
		if !(ok && token.Valid) {
			c.SetCookie(&http.Cookie{
				Name:   "access_token",
				Value:  "",
				MaxAge: -1,
			})
			return c.NoContent(http.StatusForbidden)
		}

		user, err := controllers.UserGetByUsername(claims.Username)
		if err != nil {
			c.SetCookie(&http.Cookie{
				Name:   "access_token",
				Value:  "",
				MaxAge: -1,
			})
			return c.NoContent(http.StatusForbidden)
		}

		cc := &CustomContext{c, user}
		return next(cc)
	}
}

func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)
		if !cc.IsAdmin {
			c.SetCookie(&http.Cookie{
				Name:   "access_token",
				Value:  "",
				MaxAge: -1,
			})
			return c.NoContent(http.StatusForbidden)
		}
		return next(cc)
	}
}

func CheckCustomer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)
		if cc.IsAdmin {
			c.SetCookie(&http.Cookie{
				Name:   "access_token",
				Value:  "",
				MaxAge: -1,
			})
			return c.NoContent(http.StatusForbidden)
		}
		return next(cc)
	}
}
