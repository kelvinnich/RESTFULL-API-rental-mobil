package entityservice

import (
	"restfull-api-rental-mobil/dto"
	"restfull-api-rental-mobil/models"
)

type UserService interface{
	ProfileUsers(nik string) models.User
	UpdateUsers(updateUser dto.UpdateUserDTO) models.User 
}