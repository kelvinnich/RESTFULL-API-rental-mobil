package service

import (
	"log"
	"restfull-api-rental-mobil/dto"
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/models"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)
type AuthServiceIMPL struct{
	userRepository entityrepository.UserRepository
}

func NewConnectAuthService(userRepo entityrepository.UserRepository) entityservice.AuthService{
	return &AuthServiceIMPL{
		userRepository: userRepo,
	}
}

func(service *AuthServiceIMPL) VerifyCredintial(email string, password string) interface{}{
	res := service.userRepository.VerifyCredintial(email, password)
	if k, ok := res.(models.User); ok{
		comprarePw := hashAndComparePassword(k.Password, []byte(password))
		if email ==  k.Email && comprarePw{
			return res
		}
		return false
	}
	return nil
}

func(service *AuthServiceIMPL) CreateUser(user dto.RegisterDTO) models.User{
	userRegister := models.User{}
	err := smapping.FillStruct(&userRegister, smapping.MapFields(&user) )
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	ress := service.userRepository.InsertUser(userRegister)
	return ress
}

func(service *AuthServiceIMPL) FindByWA(wa uint64) models.User{
	return service.userRepository.FindByWA(wa)
}

func(service *AuthServiceIMPL) IsDuplicateWA(wa uint64) bool{
	duplicate := service.userRepository.IsDuplicateWA(wa)
	return !(duplicate.Error == nil)
}


func hashAndComparePassword(password string, plainPW []byte) bool{
	hashPassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashPassword, plainPW)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}