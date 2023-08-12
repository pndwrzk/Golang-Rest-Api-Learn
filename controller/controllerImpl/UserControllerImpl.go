package controllerImpl

import (
	message "go-learning-restapi/constants"
	"go-learning-restapi/controller"
	"go-learning-restapi/dto"
	"go-learning-restapi/entities"
	"go-learning-restapi/services"
	"go-learning-restapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserControllerImpl(UserService services.UserService) controller.UserController{
	return &UserControllerImpl{
		UserService: UserService,
	}
}


func (u *UserControllerImpl) Registrasi(ctx *gin.Context){
  var bodyRequest entities.User
  err := ctx.ShouldBindJSON(&bodyRequest)
	if err != nil {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	get, errr := u.UserService.CreateUser(bodyRequest)
	if err != nil && err.Error() != "record not found" {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}
	if errr != nil {
		resError := dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, errr.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	resSuccess := dto.WebRespone(http.StatusOK, message.SuccessInsertData, get, message.ErrorMessageSucces)
	ctx.JSON(resSuccess.Code, resSuccess)
	return
}

func( u *UserControllerImpl)Login(ctx *gin.Context){
	 var bodyRequest entities.RequestLogin
  		err := ctx.ShouldBindJSON(&bodyRequest)
 		 if err != nil {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	_, err = u.UserService.Login(bodyRequest)
	if err != nil {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}


	token,err := utils.GenerateToken(bodyRequest.Email)
	if err != nil {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}
	refreshToken,err := utils.GenerateRefreshToken(bodyRequest.Email)
	if err != nil {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	responeLogin := dto.ResponeLogin{
	Code: http.StatusOK,
	Status: message.SuccessGetData,
	Token: token,
	RefreshToken: refreshToken,
	}

	ctx.JSON(responeLogin.Code, responeLogin)
	return

}