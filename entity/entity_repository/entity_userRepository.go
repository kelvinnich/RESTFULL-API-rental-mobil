package entityrepository

import (
	"restfull-api-rental-mobil/models"

	"gorm.io/gorm"
)

type UserRepository interface{
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredintial(email string, password string) interface{}
	IsDuplicateWA(wa uint64) (tx *gorm.DB)
	FindByWA(wa uint64) models.User
	ProfileUser(nik string) models.User
}