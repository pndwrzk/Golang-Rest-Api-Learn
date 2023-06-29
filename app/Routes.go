package app

import (
	"github.com/gin-gonic/gin"
	"go-learning-restapi/controller"
)

func Routes(controllerProduct controller.ProductController) {

	route := gin.Default()
	api := route.Group("/api")

	//route product
	api.GET("/products", controllerProduct.FindAll)
	route.Run()
}
