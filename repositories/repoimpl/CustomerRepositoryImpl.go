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
		row, err := c.DB.Raw("Select * From customers").Rows()

		if err != nil{
			return nil, err
		}

		var customers []entities.Customer
		
		for row.Next(){
			var customer entities.Customer
			err := row.Scan(
				&customer.ID,
				&customer.Address,
				&customer.Email,
				&customer.Name,
				&customer.PhoneNumber,
			)
			if err != nil{
				return nil, err
			}
			customers = append(customers, customer)
		}

		return customers, nil

}