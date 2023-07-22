package app

import (
	"go-learning-restapi/controller"

	"github.com/gin-gonic/gin"
)

func Routes(controllerProduct controller.ProductController, controllerCustomer controller.CustomerController, controllerUser controller.UserController) {

	route := gin.Default()
	api := route.Group("/api")

	//Route Product
	api.GET("/products", controllerProduct.FindAll)
	api.POST("/products", controllerProduct.Insert)

	//router Customer
	api.GET("/customers", controllerCustomer.FindAll)
	api.GET("/customers/:id", controllerCustomer.FindById)
	api.POST("/customers", controllerCustomer.Insert)
	api.PUT("/customers/:id", controllerCustomer.UpdateById)

	// router Customer 
	api.GET("/registrasi", controllerUser.Registrasi)
api.GET("/login", controllerUser.Login)
	route.Run()
}
