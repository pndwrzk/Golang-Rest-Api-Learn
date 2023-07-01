package controller

import "github.com/gin-gonic/gin"

type CustomerController interface {
	FindAll(ctx *gin.Context)
	Insert(ctx *gin.Context)
}
