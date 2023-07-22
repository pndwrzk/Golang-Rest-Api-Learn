package servicesimpl

import (
	"errors"
	"go-learning-restapi/constants"
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"
	"go-learning-restapi/services"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) services.UserService{
	return &UserServiceImpl{
		UserRepository: userRepository,
	} 
}


func (u *UserServiceImpl) CreateUser(user entities.User) (interface{}, error){


	// Check Email 
	getEmail, _ := u.UserRepository.ReadByEmail(user.Email)
	if getEmail.ID != 0 {
		return nil, errors.New(constants.ErrorEmailRegistered(user.Email))
	}

	// Hashing Password
	 hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
	
	user.Password = string(hash)
	get, err := u.UserRepository.Create(user)
		return get, err
	}
	