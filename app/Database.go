package app

import (
	"go-learning-restapi/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	dsn := "root:@(127.0.0.1:3306)/golang-rest-api-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	}
	db.AutoMigrate(&entities.Product{}, &entities.Customer{})
	return db
}
