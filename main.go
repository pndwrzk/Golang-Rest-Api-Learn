package main

import (
	"go-learning-restapi/app"
	"go-learning-restapi/controller/controllerImpl"
	"go-learning-restapi/repositories/repoimpl"
	"go-learning-restapi/services/servicesimpl"
	"log"

	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file",err.Error())
	}
}

func main() {

	

	db := app.DBConnect()
	rdb :=app.RedisConnect()
	

	// product
	productRepository := repoimpl.NewProductRepository(db)
	productServices := servicesimpl.NewProductService(productRepository)
	productController := controllerImpl.NewProductController(productServices)

	// customer
	customerRepository := repoimpl.NewCustomerRepository(db,rdb)
	customerService := servicesimpl.NewCustomerService(customerRepository)
	customerController := controllerImpl.NewCustomerController(customerService)

	// User

	userRepository := repoimpl.NewUserRepositoryImpl(db)
	userService := servicesimpl.NewUserServiceImpl(userRepository)
	userController := controllerImpl.NewUserControllerImpl(userService)
	app.Routes(productController, customerController,userController)

}
