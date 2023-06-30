package servicesimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"
	"go-learning-restapi/services"
)

type CustomerServiceImpl struct {
	CustomerRepository repositories.CustomerRepository
}

func NewCustomerService(customerService repositories.CustomerRepository) services.CustomerService{
	return &CustomerServiceImpl{
		CustomerRepository: customerService,
	}
}


func (c *CustomerServiceImpl) ViewCustomer()([]entities.Customer, error){
	get ,err := c.CustomerRepository.View()
	if err != nil{
		return nil,err
	}

	return get, nil
}