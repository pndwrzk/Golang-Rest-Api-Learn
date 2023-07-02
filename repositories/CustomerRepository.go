package repositories

import "go-learning-restapi/entities"

type CustomerRepository interface {
	Read() ([]entities.Customer, error)
	Create(customer entities.Customer) (entities.Customer, error)
	ReadByEmail(email string) (entities.Customer, error)
	ReadById(id int) (entities.Customer, error)
	UpdateById(id uint, bodyRequest entities.Customer) (interface{}, error)
}
