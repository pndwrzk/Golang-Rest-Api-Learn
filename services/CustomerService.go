package services

import (
	"go-learning-restapi/dto"
	"go-learning-restapi/entities"
)

type CustomerService interface {
	ReadCustomer(pagination dto.ResultPaginate) ([]entities.Customer, error)
	CreateCustomer(customer entities.Customer) (entities.Customer, error)
	ReadCustomerByEmail(email string) (entities.Customer, error)
	ReadCustomerById(id int) (entities.Customer, error)
	UpdateCustomerById(id int, bodyRequest entities.Customer) (interface{}, error)
}
