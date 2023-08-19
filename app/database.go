package app

import (
	"fmt"
	"go-learning-restapi/entities"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	}
	db.AutoMigrate(&entities.Product{}, &entities.Customer{}, &entities.User{})
	return db
}


func RedisConnect()  *redis.Client{
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",redisHost,redisPort),
		Password: redisPassword,
		DB: 0,
	})


return client
}
