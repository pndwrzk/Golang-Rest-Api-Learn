package repoimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"go-learning-restapi/dto"
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	DB *gorm.DB
	RDB *redis.Client
}

func NewCustomerRepository(db *gorm.DB, rdb *redis.Client) repositories.CustomerRepository {
	return &CustomerRepositoryImpl{
		DB: db,
		RDB: rdb,
	}
}

func (c *CustomerRepositoryImpl) Read(pagination dto.ResultPaginate) ([]entities.Customer,int64, error) {

	var customers []entities.Customer
	
	// err := c.DB.Order(pagination.Order).Offset(pagination.Offset).Limit(pagination.Limit).Find(&customers).Error
   var count int64
   count=0
	err:= c.DB.Table("customers").Order(pagination.Order).Offset(pagination.Offset).Limit(pagination.Limit).Find(&customers).Count(&count).Error
	if err != nil {
		return nil, 0,err
	}
	return customers, count,nil

}

func (c *CustomerRepositoryImpl) Create(customer entities.Customer) (entities.Customer, error) {
	err := c.DB.Create(&customer).Error
	return customer, err
}

func (c *CustomerRepositoryImpl) ReadByEmail(email string) (entities.Customer, error) {
	var customer entities.Customer
	err := c.DB.First(&customer, "email = ?", email).Error

	return customer, err
}

func (c *CustomerRepositoryImpl) ReadById(id int) (entities.Customer, error) {
var customer entities.Customer
	keyRedis := fmt.Sprintf("customer_id_%d",id)
	res ,err := c.RDB.Get(context.Background(),keyRedis).Result()
	if err != nil {
		err = c.DB.First(&customer, "customer_id", id).Error
		if customer.ID !=0{
			out,_:= json.Marshal(customer)
			// 1 hour
			ttl := time.Duration(1) * time.Hour
			err = c.RDB.Set(context.Background(), keyRedis, out, ttl).Err()
		}
		return customer, err
	}
	err = json.Unmarshal([]byte(res), &customer)

	return customer, err

}

func (c *CustomerRepositoryImpl) UpdateById(id uint, bodyRequest entities.Customer) (interface{}, error) {
	bodyRequest.ID = id

	err := c.DB.Save(&bodyRequest).Error

	if err != nil {
		return nil, err
	}

	return bodyRequest, nil

}
