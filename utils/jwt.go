package utils

import (
	"github.com/golang-jwt/jwt"
)

type DataClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

type CustomClaims struct {
	DataClaims
	jwt.StandardClaims
}
