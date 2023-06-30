package app

import (
	"go-learning-restapi/controller"

	"github.com/gin-gonic/gin"
)

func Routes(controllerProduct controller.ProductController, controllerCustomer controller.CustomerController) {

	route := gin.Default()
	api := route.Group("/api")

	api.GET("/products", controllerProduct.FindAll)
	api.GET("/customers", controllerCustomer.FindAll)
	route.Run()
}
