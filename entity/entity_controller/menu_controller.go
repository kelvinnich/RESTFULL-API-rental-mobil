package entitycontroller

import "github.com/gin-gonic/gin"

type MenuController interface{
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	All(ctx *gin.Context)
	FindMenuByID(ctx *gin.Context)
}