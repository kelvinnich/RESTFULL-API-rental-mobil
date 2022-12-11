package entityservice

import (
	"restfull-api-rental-mobil/dto"
	"restfull-api-rental-mobil/models"
)


type AuthService interface{
	VerifyCredintial(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByWA(wa uint64) models.User
	IsDuplicateWA(wa uint64) bool
}
