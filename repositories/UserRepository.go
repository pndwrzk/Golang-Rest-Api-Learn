package repositories

import "go-learning-restapi/entities"

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	ReadByEmail(email string)(entities.User,error)
}
