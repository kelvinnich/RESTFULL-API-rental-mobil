package entityrepository

import "restfull-api-rental-mobil/models"


type MenuRepository interface{
	InsertMenu(menu models.Menu) models.Menu
	UpdateMenu(menu models.Menu) models.Menu
	DeleteMenu(menu models.Menu)
	AllMenu() []models.Menu
	FindMenuByTipeMobil(tipeMobil string) models.Menu
	FindMenuByID(id uint64) models.Menu
}