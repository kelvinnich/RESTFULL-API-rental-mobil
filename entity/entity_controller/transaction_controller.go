package entitycontroller

import "github.com/gin-gonic/gin"

type TransactionController interface{
	InsertTx(ctx *gin.Context)
	UpdateTx(ctx *gin.Context)
	DeleteTx(ctx *gin.Context)
	AllTx(ctx *gin.Context)
}