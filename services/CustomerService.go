package services

import "go-learning-restapi/entities"

type CustomerService interface {
	ReadCustomer() ([]entities.Customer, error)
	CreateCustomer(customer entities.Customer) (entities.Customer, error)
	ReadCustomerByEmail(email string) (entities.Customer, error)
	ReadCustomerById(id int) (entities.Customer, error)
}
