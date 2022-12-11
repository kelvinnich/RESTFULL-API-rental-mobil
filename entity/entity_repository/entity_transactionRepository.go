package entityrepository

import "restfull-api-rental-mobil/models"




type TransactionRepository interface{
	InsertTransaction(trancation models.Transaction) models.Transaction
	UpdateTransaction(transaction models.Transaction) models.Transaction
	DeleteTransaction(transaction models.Transaction)
	AllTransaction() []models.Transaction
	FindTransactionByID(transactionID uint64) models.Transaction
}
