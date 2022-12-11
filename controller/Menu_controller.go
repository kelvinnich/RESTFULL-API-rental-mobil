package controller

import (
	"net/http"
	"restfull-api-rental-mobil/dto"
	entitycontroller "restfull-api-rental-mobil/entity/entity_controller"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/helper"
	"restfull-api-rental-mobil/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuControllerIMPL struct{
	menuService entityservice.MenuService
	jwtServcie entityservice.JwtService
}

func NewMenuController(ms entityservice.MenuService, js entityservice.JwtService) entitycontroller.MenuController{
	return &MenuControllerIMPL{
		menuService: ms,
		jwtServcie: js,
	}
}

func(m *MenuControllerIMPL) Insert(ctx *gin.Context){
	var menu dto.CreateMenuDTO
	err := ctx.ShouldBind(&menu)
	if err != nil {
		response := helper.ResponseERROR("failed to insert menu", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else{
		createMenu := m.menuService.InsertMenus(menu)
		response := helper.ResponseOK(true, "succes to insert menus", createMenu)
		ctx.JSON(http.StatusOK,response )
	}
}

func(m *MenuControllerIMPL) Update(ctx *gin.Context){
	var menu dto.UpdateMenuDTO
	err :=  ctx.ShouldBind(&menu)
	if err != nil {
		response := helper.ResponseERROR("sorry your update menus is failed", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"),10, 64)
	if err != nil {
		response := helper.ResponseERROR("fatal error", "please try agin ", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	menu.ID = id
	updateMenu := m.menuService.UpdateMenus(menu)
	response := helper.ResponseOK(true, "succes Update menu", updateMenu)
	ctx.JSON(http.StatusOK, response)
}

func(m *MenuControllerIMPL) Delete(ctx *gin.Context){
	var menu models.Menu
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.ResponseERROR("failed to delete", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	menu.ID = id
	m.menuService.DeleteMenus(menu)
	response := helper.ResponseOK(true, "succes delete menu", menu)
	ctx.JSON(http.StatusOK, response)
}

func(m *MenuControllerIMPL) All(ctx *gin.Context){
	var menu []models.Menu = m.menuService.AllMenus()
	response := helper.ResponseOK(true, "Succes get all menu!", menu)
	ctx.JSON(http.StatusOK, response)
}

func(m *MenuControllerIMPL) FindMenuByID(ctx *gin.Context){
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.ResponseERROR("failed to get data menu by id", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var menu models.Menu = m.menuService.FindMenusByID(id)
	if (menu == models.Menu{}) {
		response := helper.ResponseERROR("data is not compatyble", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else {
		response := helper.ResponseOK(true,"Succes to get data by id", menu)
		ctx.AbortWithStatusJSON(http.StatusOK, response)
	}
}