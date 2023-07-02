package repoimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repositories.CustomerRepository {
	return &CustomerRepositoryImpl{
		DB: db,
	}
}

func (c *CustomerRepositoryImpl) Read() ([]entities.Customer, error) {

	var customers []entities.Customer
	err := c.DB.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil

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
	err := c.DB.First(&customer, "customer_id", id).Error

	return customer, err

}
