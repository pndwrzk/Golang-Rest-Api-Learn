package controllerImpl

import (
	"go-learning-restapi/controller"
	"go-learning-restapi/dto"
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

func (p ProductControllerImpl) FindAll(ctx *gin.Context) {
	get, err := p.ProductService.ViewProduct()

	if err != nil {
		resError := dto.WebRespone(http.StatusInternalServerError, "Internal Server Error", nil,err.Error())
		ctx.JSON(resError.Code, resError)
		return
	}
	resSuccess := dto.WebRespone(http.StatusOK, "OK", get,"")
	ctx.JSON(resSuccess.Code, resSuccess)
	return
}
