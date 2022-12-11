package entityservice

import (
	"restfull-api-rental-mobil/dto"
	"restfull-api-rental-mobil/models"
)

type MenuService interface{
	InsertMenus(menu dto.CreateMenuDTO) models.Menu
	UpdateMenus(menu dto.UpdateMenuDTO) models.Menu
	DeleteMenus(menu models.Menu)
	AllMenus() []models.Menu
	FindMenusByID(menuID uint64) models.Menu
}