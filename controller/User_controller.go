package controller

import (
	"fmt"
	"net/http"
	"restfull-api-rental-mobil/dto"
	entitycontroller "restfull-api-rental-mobil/entity/entity_controller"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/helper"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserControllerIMPL struct {
	userService entityservice.UserService
	jwtService  entityservice.JwtService
}

func NewConnectUserController(us entityservice.UserService, js entityservice.JwtService) entitycontroller.UserController {
	return &UserControllerIMPL{
		userService: us,
		jwtService:  js,
	}
}

func (c *UserControllerIMPL) Profiles(context *gin.Context) {
	autheader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(autheader)
	if err != nil {
		response := helper.ResponseERROR("vailed to request profile", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	nik := fmt.Sprintf("%v", claims["nik"])
	user := c.userService.ProfileUsers(nik)
	response := helper.ResponseOK(true, "Succes Get Profile", user)
	context.JSON(http.StatusOK, response)
}

func (c *UserControllerIMPL) Updates(context *gin.Context) {
	var userUpdate dto.UpdateUserDTO
	err := context.ShouldBind(&userUpdate)
	if err != nil {
		response := helper.ResponseERROR("failed to update", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	autheader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(autheader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	nik, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	userUpdate.ID = int(nik)
	user := c.userService.UpdateUsers(userUpdate)
	response := helper.ResponseOK(true, "succes to update", user)
	context.JSON(http.StatusOK, response)
}
