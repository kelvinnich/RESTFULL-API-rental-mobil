package entitycontroller

import "github.com/gin-gonic/gin"

type AuthController interface{
	Register(context *gin.Context)
	Login(context *gin.Context)
}