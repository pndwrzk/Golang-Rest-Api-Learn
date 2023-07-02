package controller

import "github.com/gin-gonic/gin"

type CustomerController interface {
	FindAll(ctx *gin.Context)
	Insert(ctx *gin.Context)
	FindById(ctx *gin.Context)
	UpdateById(ctx *gin.Context)
}
