package app

import (
	"go-learning-restapi/controller"
	"go-learning-restapi/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(controllerProduct controller.ProductController, controllerCustomer controller.CustomerController, controllerUser controller.UserController) {

	route := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} 
    config.AllowMethods = []string{"GET", "POST", "PUT",  "DELETE"}
	config.AllowCredentials=true
	route.Use(cors.New(config))



	route.POST("/login", controllerUser.Login)
	api := route.Group("/api")

	api.Use(middleware.AuthenticateMiddleware())

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

	route.Run()
}
