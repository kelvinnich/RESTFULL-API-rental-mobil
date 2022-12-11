package service

import (
	"log"
	"restfull-api-rental-mobil/dto"
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/models"

	"github.com/mashingan/smapping"
)

type ConnectUserServiceIMPLEMNT struct{
	userService entityrepository.UserRepository
}

func NewConnectUserService(user entityrepository.UserRepository)entityservice.UserService{

	return &ConnectUserServiceIMPLEMNT{
		userService: user,
	}
}

func(service *ConnectUserServiceIMPLEMNT) ProfileUsers(nik string) models.User{
return service.userService.ProfileUser(nik)
}

func(service *ConnectUserServiceIMPLEMNT) UpdateUsers(updateUser dto.UpdateUserDTO) models.User{
	userUpdate := models.User{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&updateUser))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}

	ress := service.userService.UpdateUser(userUpdate)
	return ress
}