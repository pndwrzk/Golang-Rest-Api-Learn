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

	var respone dto.Respone
	get, err := p.ProductService.ReadProduct()

	if err != nil {
		respone = dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, err.Error())
		ctx.JSON(respone.Code, respone)
		return
	}
	respone = dto.WebRespone(http.StatusOK, message.SuccessGetData, get, message.ErrorMessageSucces)
	ctx.JSON(respone.Code, respone)
	return
}

func (p *ProductControllerImpl) Insert(ctx *gin.Context) {
	var bodyRequest entities.Product
	var respone dto.Respone
	err := ctx.ShouldBindJSON(&bodyRequest)
	if err != nil {
		respone = dto.WebRespone(http.StatusNotFound, message.ErrorStatus, nil, err.Error())
		ctx.JSON(respone.Code, respone)
		return
	}

	get, errr := p.ProductService.CreateProduct(bodyRequest)
	if errr != nil {
		respone = dto.WebRespone(http.StatusInternalServerError, message.ErrorStatus, nil, errr.Error())
		ctx.JSON(respone.Code, respone)
		return
	}

	respone = dto.WebRespone(http.StatusOK, message.SuccessInsertData, get, message.ErrorMessageSucces)
	ctx.JSON(respone.Code, respone)
	return

}
