package services

import "go-learning-restapi/entities"

type CustomerService interface {
	ViewCustomer() ([]entities.Customer, error)
}
