package services

import "go-learning-restapi/entities"

type CustomerService interface {
	ReadCustomer() ([]entities.Customer, error)
	CreateCustomer(customer entities.Customer) (entities.Customer, error)
}
