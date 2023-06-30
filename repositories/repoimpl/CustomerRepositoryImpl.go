package repoimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	DB *gorm.DB
}


func NewCustomerRepository(db *gorm.DB) repositories.CustomerRepository{
return &CustomerRepositoryImpl{
	DB: db,
}
}


func (c *CustomerRepositoryImpl) View()([]entities.Customer, error){
		
		var customers []entities.Customer
		err := c.DB.Find(&customers).Error
			if err != nil{
				return nil,err
			}
		return customers, nil

}