package controllerImpl

import (
	message "go-learning-restapi/constants"
	"go-learning-restapi/controller"
	"go-learning-restapi/dto"
	"go-learning-restapi/entities"
	"go-learning-restapi/services"
	"go-learning-restapi/utils"
	"net/http"
	"strconv"

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
	var respone dto.Respone
	var responeGetAll dto.ResponeGetAll
	pagination,pagging := utils.GeneratePagination(ctx)
	get,totalData, err := c.CustomerService.ReadCustomer(pagination)

	if err != nil {
		respone = dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, err.Error())
		ctx.JSON(respone.Code, respone)
		return
	}

	responeGetAll = dto.WebResponeGetAll(http.StatusOK, message.SuccessGetData, get,totalData,pagging, message.ErrorMessageSucces)
	ctx.JSON(responeGetAll.Code, responeGetAll)
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

	//check customer by email
	get, err := c.CustomerService.ReadCustomerByEmail(bodyRequest.Email)
	if err != nil && err.Error() != "record not found" {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	if get.ID != 0 {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, message.ErrorEmailRegistered(get.Email))
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

func (c *CustomerControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	var respone dto.Respone
	idInt, err := strconv.Atoi(id)


	if err != nil {
		respone = dto.WebRespone(http.StatusBadRequest, message.ErrorStatus, nil, message.ErrorIdNotValid)
		ctx.JSON(respone.Code, respone)
		return
	}

	res, err := c.CustomerService.ReadCustomerById(idInt)

	if err != nil && err.Error() != "record not found" {
		respone = dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, err.Error())
		ctx.JSON(respone.Code, respone)
		return
	}

	if res.ID == 0 {
		 respone = dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, message.ErrorDataNotFound)
		ctx.JSON(respone.Code, respone)
		return
	}

	respone = dto.WebRespone(http.StatusOK, message.SuccessGetData, res, message.ErrorMessageSucces)
	ctx.JSON(respone.Code, respone)
	return

}

func (c *CustomerControllerImpl) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")

	var respone dto.Respone

	idInt, err := strconv.Atoi(id)
	if err != nil {
		respone = dto.WebRespone(http.StatusBadRequest, message.ErrorStatus, nil, message.ErrorIdNotValid)
		ctx.JSON(respone.Code, respone)
		return
	}

	var bodyRequest entities.Customer
	err = ctx.ShouldBindJSON(&bodyRequest)
	if err != nil {
		respone = dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(respone.Code, respone)
		return
	}

	res, err := c.CustomerService.UpdateCustomerById(idInt, bodyRequest)
	if err != nil {
		respone = dto.WebRespone(http.StatusBadRequest, message.ErrorStatus, nil, err.Error())
		ctx.JSON(respone.Code, respone)
		return
	}

	respone = dto.WebRespone(http.StatusOK, message.SuccessUpdateDate, res, message.ErrorMessageSucces)
	ctx.JSON(respone.Code, respone)
	return

}
