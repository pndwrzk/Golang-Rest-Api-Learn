package servicesimpl

import (
	"errors"
	"go-learning-restapi/constants"
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

func (c *CustomerServiceImpl) ReadCustomerById(id int) (entities.Customer, error) {
	get, err := c.CustomerRepository.ReadById(id)
	return get, err
}

func (c *CustomerServiceImpl) UpdateCustomerById(id int, bodyRequest entities.Customer) (interface{}, error) {

	getId, err := c.CustomerRepository.ReadById(id)
	if err != nil {
		return nil, err
	}

	getEmail, _ := c.CustomerRepository.ReadByEmail(bodyRequest.Email)

	if getEmail.ID != 0 && bodyRequest.Email != getId.Email {
		return nil, errors.New(constants.ErrorEmailRegistered(bodyRequest.Email))
	}

	idUint := uint(id)
	res, err := c.CustomerRepository.UpdateById(idUint, bodyRequest)
	if err != nil {
		return nil, err
	}

	return res, err

}
