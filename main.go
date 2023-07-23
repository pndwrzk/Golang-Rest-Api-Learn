package main

import (
	"context"
	"fmt"
	"go-learning-restapi/app"
	"go-learning-restapi/controller/controllerImpl"
	"go-learning-restapi/repositories/repoimpl"
	"go-learning-restapi/services/servicesimpl"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file",err.Error())
	}

	db := app.DBConnect()
	rdb :=app.RedisConnect()
	key:= "key-123"
	 // get data
    op2 := rdb.Get(context.Background(),key)
    if err := op2.Err(); err != nil {
        fmt.Printf("unable to GET data. error: %v", err)
        return
    }
    res, err := op2.Result()
    if err != nil {
        fmt.Printf("unable to GET data. error: %v", err)
        return
    }
    log.Println("get operation success. result:", res)

	// product
	productRepository := repoimpl.NewProductRepository(db)
	productServices := servicesimpl.NewProductService(productRepository)
	productController := controllerImpl.NewProductController(productServices)

	// customer
	customerRepository := repoimpl.NewCustomerRepository(db)
	customerService := servicesimpl.NewCustomerService(customerRepository)
	customerController := controllerImpl.NewCustomerController(customerService)

	// User

	userRepository := repoimpl.NewUserRepositoryImpl(db)
	userService := servicesimpl.NewUserServiceImpl(userRepository)
	userController := controllerImpl.NewUserControllerImpl(userService)
	app.Routes(productController, customerController,userController)

}
