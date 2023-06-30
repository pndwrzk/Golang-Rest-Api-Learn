package controllerImpl

import (
	"go-learning-restapi/controller"
	"go-learning-restapi/dto"
	"go-learning-restapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerControllerImpl struct {
	CustomerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) controller.CustomerController{
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func(c *CustomerControllerImpl) FindAll(ctx *gin.Context){
  get, err := c.CustomerService.ViewCustomer()
  if err != nil{
	resError := dto.WebRespone(http.StatusInternalServerError, "Internal Server Error", nil,err.Error())
		ctx.JSON(resError.Code, resError)
		return
  }

 resSuccess := dto.WebRespone(http.StatusOK, "OK", get,"")
ctx.JSON(resSuccess.Code, resSuccess)
return
	
}