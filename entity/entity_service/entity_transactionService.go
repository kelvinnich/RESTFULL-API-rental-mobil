package entityservice

import (
	"restfull-api-rental-mobil/dto"
	"restfull-api-rental-mobil/models"
)

type TransactionService interface{
	InsertTransactions(t dto.CreateTransactionDTO) models.Transaction
	UpdateTransactions(t dto.UpdateTransactionDTO) models.Transaction
	Deletetransactions(t models.Transaction)
	AllTransactions() []models.Transaction
	FindtransactonsById(txId uint64) models.Transaction
	IsAllowedtoEditTransactions(userID string, txId uint64) bool
}