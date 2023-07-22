package services

import "go-learning-restapi/entities"

type UserService interface {
	CreateUser(user entities.User) (interface{}, error)
	Login(user entities.RequestLogin)(interface{},error)
}
