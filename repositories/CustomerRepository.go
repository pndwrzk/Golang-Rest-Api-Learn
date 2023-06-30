package repositories

import "go-learning-restapi/entities"

type CustomerRepository interface {
	View() ([]entities.Customer, error)
}
