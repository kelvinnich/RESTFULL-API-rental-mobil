package entitycontroller

import "github.com/gin-gonic/gin"

type UserController interface{
	Profiles(context *gin.Context)
	Updates(context *gin.Context)
}