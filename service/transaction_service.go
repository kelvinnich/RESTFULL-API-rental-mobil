package service

import (
	"fmt"
	"restfull-api-rental-mobil/dto"
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/models"
)

type ConnecTransactionServiceIMPL struct{
	transactionRepository entityrepository.TransactionRepository
	menuRepository entityrepository.MenuRepository
}

func NewConnectTransactionsService(txRepo entityrepository.TransactionRepository, menuRepo entityrepository.MenuRepository) entityservice.TransactionService{
	return &ConnecTransactionServiceIMPL{
		transactionRepository: txRepo,
		menuRepository: menuRepo,
	}
}

func(service *ConnecTransactionServiceIMPL) InsertTransactions(t dto.CreateTransactionDTO)models.Transaction{
	tx := models.Transaction{}
	tx.UserID = t.UserID
	tx.MenuID = tx.MenuID
	tx.DurasaRental = uint64(t.DurasiPinjam)
	tx.TotalPembayaran = service.menuRepository.FindMenuByID(t.MenuID).Harga * uint64(t.DurasiPinjam)

	res := service.transactionRepository.InsertTransaction(tx)
	return res
}

func(service *ConnecTransactionServiceIMPL) UpdateTransactions(t dto.UpdateTransactionDTO)models.Transaction{
	tx := models.Transaction{}
	tx.ID = uint64(t.ID)
	tx.UserID = uint64(t.UserID)
	tx.MenuID = tx.MenuID
	tx.DurasaRental = uint64(t.DurasiPinjam)
	tx.TotalPembayaran = service.menuRepository.FindMenuByID(t.MenuID).Harga * uint64(t.DurasiPinjam)

	res := service.transactionRepository.UpdateTransaction(tx)
	return res
}

func(service *ConnecTransactionServiceIMPL) Deletetransactions(t models.Transaction){
	service.transactionRepository.DeleteTransaction(t)
}

func(service *ConnecTransactionServiceIMPL) AllTransactions() []models.Transaction{
	return service.transactionRepository.AllTransaction()
}

func(service *ConnecTransactionServiceIMPL) FindtransactonsById(txId uint64) models.Transaction{
	return service.transactionRepository.FindTransactionByID(txId)
}

func(service *ConnecTransactionServiceIMPL) IsAllowedtoEditTransactions(userID string, txId uint64) bool{
	tx := service.transactionRepository.FindTransactionByID(txId)
	id := fmt.Sprintf("%v", tx.UserID)
	return userID == id
}