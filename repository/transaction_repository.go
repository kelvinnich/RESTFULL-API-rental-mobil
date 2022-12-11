package repository

import (
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	"restfull-api-rental-mobil/models"

	"gorm.io/gorm"
)

type ConnectTransacation struct{
connected *gorm.DB
}

func ConnecetTransactionRepository(db *gorm.DB) entityrepository.TransactionRepository{
	return &ConnectTransacation{
		connected: db,
	}
}

func(db *ConnectTransacation) InsertTransaction(tx models.Transaction) models.Transaction{
	db.connected.Exec("INSERT INTO transactions (user_id, menu_id, durasi_rental, total_pembayaran, tanggal) VALUES (?,?,?,?,?)", tx.UserID, tx.MenuID, tx.DurasaRental, tx.TotalPembayaran, tx.Tanggal)
	return tx
}

func(db *ConnectTransacation) UpdateTransaction(tx models.Transaction) models.Transaction{
	db.connected.Exec("UPDATE transactions SET user_id = ?, menu_id = ?, durasi_rental = ?, total_pembayaran = ?, tanggal = ?, WHERE id = ?", tx.UserID, tx.MenuID, tx.DurasaRental, tx.TotalPembayaran, tx.Tanggal, tx.ID)
	return tx
}

func(db *ConnectTransacation) DeleteTransaction(tx models.Transaction){
	db.connected.Delete(tx)
}

func(db *ConnectTransacation) AllTransaction() []models.Transaction{
	var tx []models.Transaction
	db.connected.Preload("User").Find(&tx)
	return tx
}

func(db *ConnectTransacation) FindTransactionByID(txID uint64) models.Transaction{
	var tx models.Transaction
	db.connected.Find(&tx, txID)
	return tx
}