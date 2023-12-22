package middleware

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
