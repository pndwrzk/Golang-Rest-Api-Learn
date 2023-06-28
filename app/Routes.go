package app

import (
	"go-learning-restapi/controller"

	"github.com/gin-gonic/gin"
)


func Routes(controlerProduct *controller.ProductContrller) {
	r := gin.Default()
	r.GET("/products", controlerProduct.FindAll)
	r.Run()
}
