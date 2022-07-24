package utils

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}
