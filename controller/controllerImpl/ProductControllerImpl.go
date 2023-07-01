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

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) controller.ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (p *ProductControllerImpl) FindAll(ctx *gin.Context) {
	get, err := p.ProductService.ViewProduct()

	if err != nil {
		resError := dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}
	resSuccess := dto.WebRespone(http.StatusOK, message.SuccessStatus, get, message.ErrorMessageSucces)
	ctx.JSON(resSuccess.Code, resSuccess)
	return
}

func (p *ProductControllerImpl) Create(ctx *gin.Context) {
	var bodyRequest entities.Product
	err := ctx.ShouldBindJSON(&bodyRequest)
	if err != nil {
		resError := dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	get, errr := p.ProductService.CreateProduct(bodyRequest)
	if errr != nil {
		resError := dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, errr.Error())
		ctx.JSON(resError.Code, resError)
		return
	}

	resSuccess := dto.WebRespone(http.StatusOK, message.SuccessInsertData, get, message.ErrorMessageSucces)
	ctx.JSON(resSuccess.Code, resSuccess)
	return

}
