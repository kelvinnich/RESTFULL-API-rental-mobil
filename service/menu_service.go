package service

import (
	"log"
	"restfull-api-rental-mobil/dto"
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/models"

	"github.com/mashingan/smapping"
)

type MenuServiceIMPL struct{
	menuRepository entityrepository.MenuRepository
}

func NewConnectMenuService(m entityrepository.MenuRepository) entityservice.MenuService{
	return &MenuServiceIMPL{
		menuRepository: m,
	}
}

func(service *MenuServiceIMPL) InsertMenus(m dto.CreateMenuDTO) models.Menu{
	 menu := models.Menu{}
	err := smapping.FillStruct(&menu, smapping.MapFields(&m))
	if err != nil {
		log.Println(err)
		panic("failed to add menu")
	}
	ress := service.menuRepository.InsertMenu(menu)
	return ress
}

func(service *MenuServiceIMPL) UpdateMenus(m dto.UpdateMenuDTO) models.Menu{
	var menu models.Menu
	err := smapping.FillStruct(&menu, smapping.MapFields(&m))
	if err != nil {
		log.Println(err)
		panic("failed to add menu")
	}

	ress := service.menuRepository.UpdateMenu(menu)
	return ress
}

func(service *MenuServiceIMPL) DeleteMenus(m models.Menu){
	service.menuRepository.DeleteMenu(m)
}

func(service *MenuServiceIMPL) AllMenus() []models.Menu{
	return service.menuRepository.AllMenu()
}

func(service *MenuServiceIMPL) FindMenusByID(menuID uint64) models.Menu{
	
	return service.menuRepository.FindMenuByID(menuID)
}
