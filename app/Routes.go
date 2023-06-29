package app

import (
	"github.com/gin-gonic/gin"
	"go-learning-restapi/controller"
)

func Routes(controlerProduct *controller.ProductController) {
	r := gin.Default()
	r.GET("/products", controlerProduct.FindAll)
	r.Run()
}
