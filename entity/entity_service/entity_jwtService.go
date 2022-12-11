package entityservice

import "github.com/dgrijalva/jwt-go"

type JwtService interface{
	GenerateToken(userId uint64) string
	ValidateToken(token string) (*jwt.Token, error)
}