package controllerImpl

import (
	"github.com/gin-gonic/gin"
	"go-learning-restapi/controller"
	"go-learning-restapi/dto"
	"go-learning-restapi/services"
	"net/http"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) controller.ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (p ProductControllerImpl) FindAll(ctx *gin.Context) {
	get, err := p.ProductService.ViewProduct()

	if err != nil {
		resError := dto.WebResponeError(http.StatusInternalServerError, "Error", err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}
	resSuccess := dto.WebResponeSuccess(http.StatusOK, "OK", get)
	ctx.JSON(resSuccess.Code, resSuccess)
	return
}
