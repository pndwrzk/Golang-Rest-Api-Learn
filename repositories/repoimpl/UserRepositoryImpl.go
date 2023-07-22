package repoimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repositories.UserRepository{
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (u *UserRepositoryImpl) Create(user entities.User)(entities.User, error){
	err := u.DB.Create(&user).Error
	return user, err

}
func (u *UserRepositoryImpl) ReadByEmail(email string) (entities.User, error) {
	var user entities.User
	err := u.DB.First(&user, "email = ?", email).Error
	return user, err
}