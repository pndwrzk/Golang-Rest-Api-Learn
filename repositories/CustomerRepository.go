package repositories

import "go-learning-restapi/entities"

type CustomerRepository interface {
	Read() ([]entities.Customer, error)
	Create(customer entities.Customer) (entities.Customer, error)
}
