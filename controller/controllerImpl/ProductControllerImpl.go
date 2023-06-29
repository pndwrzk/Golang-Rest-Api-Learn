package controllerImpl

import (
	"github.com/gin-gonic/gin"
	"go-learning-restapi/controller"
	"go-learning-restapi/services"
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
		ctx.AbortWithStatusJSON(400, err)
	}

	ctx.JSON(200, get)
}
