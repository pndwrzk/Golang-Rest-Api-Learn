package servicesimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"
	"go-learning-restapi/services"
)

type CustomerServiceImpl struct {
	CustomerRepository repositories.CustomerRepository
}

func NewCustomerService(customerService repositories.CustomerRepository) services.CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerService,
	}
}

func (c *CustomerServiceImpl) ReadCustomer() ([]entities.Customer, error) {
	get, err := c.CustomerRepository.Read()
	if err != nil {
		return nil, err
	}

	return get, nil
}

func (c *CustomerServiceImpl) CreateCustomer(customer entities.Customer) (entities.Customer, error) {
	get, err := c.CustomerRepository.Create(customer)
	return get, err
}

func (c *CustomerServiceImpl) ReadCustomerByEmail(email string) (entities.Customer, error) {
	get, err := c.CustomerRepository.ReadByEmail(email)
	return get, err
}
