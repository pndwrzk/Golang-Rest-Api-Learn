package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Registrasi(ctx *gin.Context)
	Login(ctx *gin.Context)
}
