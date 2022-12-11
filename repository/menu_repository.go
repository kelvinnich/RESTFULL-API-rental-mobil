package repository

import (
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	"restfull-api-rental-mobil/models"

	"gorm.io/gorm"
)
type MenuConnection struct{
	connected *gorm.DB
}

func ConnectMenuRepository(db *gorm.DB) entityrepository.MenuRepository{
	return &MenuConnection{
		connected: db,
	}
}

func(db *MenuConnection) InsertMenu(menu models.Menu) models.Menu{
	db.connected.Exec("INSERT INTO menus(id, nama_mobil, harga, tipe_mobil, status) VALUES (?,?,?,?,?)", menu.ID, menu.NamaMobil, menu.Harga, menu.TipeMobil, menu.Status)
	return menu
}

func(db *MenuConnection) UpdateMenu(menu models.Menu) models.Menu{
	db.connected.Exec("UPDATE menus SET nama_mobil = ?, harga = ?, tipe_mobil = ?, status = ? WHERE id = ?", menu.NamaMobil, menu.Harga, menu.TipeMobil, menu.Status, menu.ID)
	return menu
}

func(db *MenuConnection) DeleteMenu(menu models.Menu){
	db.connected.Delete(menu)
}

func(db *MenuConnection) AllMenu() []models.Menu{
	var menu []models.Menu
	db.connected.Find(&menu)
	return menu
}

func(db *MenuConnection) FindMenuByTipeMobil(tipeMobil string) models.Menu{
	var menu models.Menu
	db.connected.Where("tipe_mobil = ?", tipeMobil).Take(&menu)
	return menu
}

func(db *MenuConnection) FindMenuByID(id uint64) models.Menu{
	var menu models.Menu
	db.connected.Find(&menu, id)
	return menu
}
