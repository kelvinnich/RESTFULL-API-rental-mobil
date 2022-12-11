package service

import (
	"fmt"
	"log"
	"os"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JwtCustomClaim struct{
	NIK uint64 `json:"nik"`
	jwt.StandardClaims
}

type JwtservicesIMPL struct{
	secretKey string
	issuer string
}

func NewJwtService() entityservice.JwtService{
	return &JwtservicesIMPL{
		secretKey: SecretkeyEnv(),
		issuer: "kelvinnnnnnn",
	}
}

func SecretkeyEnv() string{

	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		panic("failed to load env file")
	}
	
	secretkey := os.Getenv("JWT_SECRET")
	if secretkey != "" {
		secretkey = "kelvin"
	}

	return secretkey
}

func(j *JwtservicesIMPL) GenerateToken(nik uint64) string{
	claims := JwtCustomClaim{
		nik,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0,7,0).Unix(),
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.issuer))
	if err != nil {
		log.Println(err)
		panic("failed to generate Token")
	}
	return t
}

func(j *JwtservicesIMPL) ValidateToken(token string) (*jwt.Token, error){
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}