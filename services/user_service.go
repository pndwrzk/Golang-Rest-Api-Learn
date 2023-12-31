package services

import (
	"go-learning-restapi/dto"
	"go-learning-restapi/entities"
)

type UserService interface {
	CreateUser(user entities.User) (interface{}, error)
	Login(user entities.RequestLogin)(dto.ResponeLogin,error)
	RefreshToken(refreshToken string)(dto.ResponeLogin,error)
}
