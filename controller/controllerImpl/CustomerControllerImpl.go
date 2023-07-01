package controllerImpl

import (
	message "go-learning-restapi/constants"
	"go-learning-restapi/controller"
	"go-learning-restapi/dto"
	"go-learning-restapi/entities"
	"go-learning-restapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerControllerImpl struct {
	CustomerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) controller.CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (c *CustomerControllerImpl) FindAll(ctx *gin.Context) {
	get, err := c.CustomerService.ReadCustomer()
	if err != nil {
		resError := dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	resSuccess := dto.WebRespone(http.StatusOK, message.SuccessGetData, get, message.ErrorMessageSucces)
	ctx.JSON(resSuccess.Code, resSuccess)
	return

}

func (c *CustomerControllerImpl) Insert(ctx *gin.Context) {
	var bodyRequest entities.Customer
	err := ctx.ShouldBindJSON(&bodyRequest)
	if err != nil {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	get, errr := c.CustomerService.CreateCustomer(bodyRequest)
	if errr != nil {
		resError := dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, errr.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	resSuccess := dto.WebRespone(http.StatusOK, message.SuccessInsertData, get, message.ErrorMessageSucces)
	ctx.JSON(resSuccess.Code, resSuccess)
	return
}
