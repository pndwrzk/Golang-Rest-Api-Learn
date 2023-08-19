package servicesimpl

import (
	"errors"
	"go-learning-restapi/constants"
	"go-learning-restapi/dto"
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"
	"go-learning-restapi/services"
	"go-learning-restapi/utils"
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
func (u *UserServiceImpl) Login(user entities.RequestLogin)(dto.ResponeLogin,error){
	// get user by email
	get, _ := u.UserRepository.ReadByEmail(user.Email)
	if get.ID == 0 {
		return dto.ResponeLogin{}, errors.New(constants.ErrorFailedLogin)
	}

	// check password
	err := bcrypt.CompareHashAndPassword([]byte(get.Password), []byte(user.Password))
   if err != nil {
		return dto.ResponeLogin{}, errors.New(constants.ErrorFailedLogin)
	}

	// create token

	token,timerToken,err := utils.GenerateToken(get.ID)
	if err != nil {
		return dto.ResponeLogin{}, err
	}
	refreshToken,timeRefreshToken,err := utils.GenerateRefreshToken(get.ID)
	if err != nil {
		return dto.ResponeLogin{}, err
	}

	resp := dto.ResponeLogin{
	Token: token,
	RefreshToken: refreshToken,
	TokenExpired: timerToken,
	RefreshTokenExpired: timeRefreshToken,
	}

	return resp, nil



}

func (u *UserServiceImpl) RefreshToken(refreshToken string)(dto.ResponeLogin,error){

	

	idUser, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return dto.ResponeLogin{}, err
	}

	token,timerToken,err := utils.GenerateToken(uint(idUser))
	if err != nil {
		return dto.ResponeLogin{}, err
	}
	refreshToken,timeRefreshToken,err := utils.GenerateRefreshToken(uint(idUser))
	if err != nil {
		return dto.ResponeLogin{}, err
	}

	resp := dto.ResponeLogin{
	Token: token,
	RefreshToken: refreshToken,
	TokenExpired: timerToken,
	RefreshTokenExpired: timeRefreshToken,
	}

	return resp, nil

}

	