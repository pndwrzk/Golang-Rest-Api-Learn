package main

import (
	"go-learning-restapi/app"
	"go-learning-restapi/controller/controllerImpl"
	"go-learning-restapi/repositories/repoimpl"
	"go-learning-restapi/services/servicesimpl"
)

func main() {

	db := app.DBConnect()
	productRepository := repoimpl.NewProductRepository(db)
	productServices := servicesimpl.NewProductService(productRepository)
	productController := controllerImpl.NewProductController(productServices)
	app.Routes(productController)

}
