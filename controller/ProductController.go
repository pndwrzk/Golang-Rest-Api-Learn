package controller

import (
	"go-learning-restapi/services"

	"github.com/gin-gonic/gin"
)

type ProductContrller struct {
	productService services.ProductService
}


func NewProductController(productService services.ProductService) *ProductContrller{
			return &ProductContrller{
				productService: productService,
			}
}

func (p *ProductContrller,) FindAll(c *gin.Context) {	
	get, err := p.productService.ViewProduct()
	if err != nil{
		c.AbortWithError(500,err)
	}
	c.AbortWithStatusJSON(200,get)
	
}