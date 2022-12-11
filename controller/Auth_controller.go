package controller

import (
	"net/http"
	"restfull-api-rental-mobil/dto"
	entitycontroller "restfull-api-rental-mobil/entity/entity_controller"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/helper"
	"restfull-api-rental-mobil/models"
	
	"github.com/gin-gonic/gin"
)

type ConnectAuthControllerIMPL struct {
	authService entityservice.AuthService
	jwtService  entityservice.JwtService
}

func NewVerifyController(as entityservice.AuthService, js entityservice.JwtService) entitycontroller.AuthController {
	return &ConnectAuthControllerIMPL{
		authService: as,
		jwtService:  js,
	}
}

func (j *ConnectAuthControllerIMPL) Register(context *gin.Context) {
	var dtoRegister dto.RegisterDTO
	err := context.ShouldBind(&dtoRegister)
	if err != nil {
		response := helper.ResponseERROR("failed request please check your credintial", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !j.authService.IsDuplicateWA(dtoRegister.NoWa) {
		response := helper.ResponseERROR("failed to request duplicate wa", "wa duplicate", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusConflict, response)
		return
	} else {
		createUser := j.authService.CreateUser(dtoRegister)
		token := j.jwtService.GenerateToken(createUser.NIK)
		createUser.Token = token
		response := helper.ResponseOK(true, "succes Register account, welcome !", createUser)
		context.JSON(http.StatusOK, response)
	}
}

func (j *ConnectAuthControllerIMPL) Login(context *gin.Context) {
	var loginDto dto.LoginDTO
	err := context.ShouldBind(&loginDto)
	if err != nil {
		response := helper.ResponseERROR("failed to login", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result := j.authService.VerifyCredintial(loginDto.Email, loginDto.Password)
	if v, ok := result.(models.User); ok {
		generateToken := j.jwtService.GenerateToken(v.NIK)
		v.Token = generateToken
		response := helper.ResponseOK(true, "Succes login!", v)
		context.JSON(http.StatusOK, response)
		return
	}
	response := helper.ResponseERROR("please check agin yout credintial", "check email, password, and no wa ", helper.EmptyObj{})
	context.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
