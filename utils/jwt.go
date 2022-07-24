package utils

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	IsAdmin    bool   `json:"is_admin"`
	IsVerified bool   `json:"is_verified"`
	jwt.StandardClaims
}
