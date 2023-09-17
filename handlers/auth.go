package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kristabdi/bnmo-backend/controllers"
	"github.com/kristabdi/bnmo-backend/models"
	"github.com/kristabdi/bnmo-backend/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	user := new(models.User)
	var err error

	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var dbUser models.User
	dbUser, err = controllers.UserGetByUsername(user.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if !dbUser.IsVerified {
		return echo.NewHTTPError(http.StatusUnauthorized, "User is not verified")
	}

	expiry := time.Now().Add(time.Hour * 24)

	claims := utils.CustomClaims{
		DataClaims: utils.DataClaims{
			Username: dbUser.Username,
			IsAdmin:  dbUser.IsAdmin,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err
	}

	response := struct {
		Token    string `json:"access_token"`
		Username string `json:"username"`
		Name     string `json:"name"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Token:    tokenSigned,
		Username: dbUser.Username,
		Name:     dbUser.Name,
		IsAdmin:  dbUser.IsAdmin,
	}

	return c.JSON(http.StatusOK, response)
}

func Registration(c echo.Context) error {
	user := new(models.User)
	log.Println(c.FormValue("username"))
	user.Username = c.FormValue("username")
	log.Println(user.Username)
	user.Name = c.FormValue("name")
	log.Println(c.FormValue("name"))
	log.Println(user.Name)
	log.Println(c.FormValue("password"))

	hashed, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), 14)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("hashing error"))
	}

	user.Password = string(hashed)
	log.Println(user.Password)
	photo, err := c.FormFile("photo")
	src, err := photo.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	dst, err := os.Create(filepath.Join("files", filepath.Base(base64.URLEncoding.EncodeToString([]byte(photo.Filename)))))
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	user.Photo = filepath.Base(base64.URLEncoding.EncodeToString([]byte(photo.Filename)))
	log.Println(user.Photo)
	if err := controllers.UserInsertOne(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusCreated)
}
