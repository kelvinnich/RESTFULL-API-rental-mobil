package controller

import (
	"fmt"
	"net/http"
	"restfull-api-rental-mobil/dto"
	entitycontroller "restfull-api-rental-mobil/entity/entity_controller"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	"restfull-api-rental-mobil/helper"
	"restfull-api-rental-mobil/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionControllerIMPL struct{
	transactionService entityservice.TransactionService
	jwtService entityservice.JwtService
}

func NewTransactionController(ts entityservice.TransactionService, js entityservice.JwtService) entitycontroller.TransactionController{
	return &TransactionControllerIMPL{
		transactionService: ts,
		jwtService: js,
	}
}

func(t *TransactionControllerIMPL) InsertTx(ctx *gin.Context){
	var createTX dto.CreateTransactionDTO
	err := ctx.ShouldBind(&createTX)
	if err != nil {
		response := helper.ResponseERROR("failed to insert transaction", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else{
		token := ctx.GetHeader("Authorization")
		id := t.GetCustomerIdByToken(token)
		convertedId, err := strconv.ParseUint(id, 0, 0)
		if err == nil {
			createTX.UserID = convertedId
		}
		result := t.transactionService.InsertTransactions(createTX)
		response := helper.ResponseOK(true, "Succes insert transaction", result)
		ctx.JSON(http.StatusOK, response)
	}
}

func(t *TransactionControllerIMPL) UpdateTx(ctx *gin.Context){
	var updateTx dto.UpdateTransactionDTO
	err := ctx.ShouldBind(&updateTx)
	if err != nil {
		response := helper.ResponseERROR("failed to insert transaction", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else{
		token := ctx.GetHeader("Authorization")
		id := t.GetCustomerIdByToken(token)
		convertedId, err := strconv.ParseUint(id, 0, 0)
		if err == nil {
			updateTx.UserID = convertedId
		}
		result := t.transactionService.UpdateTransactions(updateTx)
		response := helper.ResponseOK(true, "Succes update transaction", result)
		ctx.JSON(http.StatusOK, response)
	}
}

func(t *TransactionControllerIMPL) DeleteTx(ctx *gin.Context){
	var tx models.Transaction
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ResponseERROR("failed to delete", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else{
		tx.ID = id
		t.transactionService.Deletetransactions(tx)
		response := helper.ResponseOK(true, "succes delete transaction", tx)
		ctx.JSON(http.StatusOK, response)
	}
}

func(t *TransactionControllerIMPL) AllTx(ctx *gin.Context){
	var tx []models.Transaction = t.transactionService.AllTransactions()
	response := helper.ResponseOK(true, "succes get all data transaction", tx)
	ctx.JSON(http.StatusOK, response)
}

func (t *TransactionControllerIMPL) GetCustomerIdByToken(token string) string{
	tkn, err := t.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := tkn.Claims.(jwt.MapClaims)
	nik := fmt.Sprintf("%v", claims["nik"])
	return nik
}