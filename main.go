package main

import (
	"go-learning-restapi/app"
	"go-learning-restapi/controller/controllerImpl"
	"go-learning-restapi/repositories/repoimpl"
	"go-learning-restapi/services/servicesimpl"
)

func main() {

	db := app.DBConnect()

	// product
	productRepository := repoimpl.NewProductRepository(db)
	productServices := servicesimpl.NewProductService(productRepository)
	productController := controllerImpl.NewProductController(productServices)


	// customer
	customerRepository := repoimpl.NewCustomerRepository(db)
	customerService := servicesimpl.NewCustomerService(customerRepository)
	customerController := controllerImpl.NewCustomerController(customerService)
	app.Routes(productController,customerController)

}
