package utils

import (
	"fmt"
	"go-learning-restapi/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paginate struct {
	Limit    int  `json:"limit"`
	OrderBy  string  `json:"order_by"`
	OrderDir string  `json:"order_dir"`
	Offset   int  `json:"offset"`
}

func GeneratePagination(ctx *gin.Context)(dto.ResultPaginate, Paginate){
	Pagging := Paginate{}
	perPage := ctx.Query("limit")
	orderBy := ctx.Query("order_by")
	orderDir:= ctx.Query("order_dir")
	page := ctx.Query("page")

	if perPage != ""{
		intVar, err := strconv.Atoi(perPage)
		if err != nil{
			Pagging.Limit = 10
		}else{
			Pagging.Limit = intVar
		}
	}else{
		Pagging.Limit = 10
	}

	if page != ""{
		intVar, err := strconv.Atoi(page)
		if err != nil{
			Pagging.Offset = GenerateOffSate(1,Pagging.Limit)
		}else{
			Pagging.Offset = GenerateOffSate(intVar,Pagging.Limit)
		}
	}else{
		Pagging.Offset = GenerateOffSate(1,Pagging.Limit)
	}

	if orderBy != ""{
		Pagging.OrderBy = orderBy
	}else{
		Pagging.OrderBy = "customer_id"
	}


	if orderDir != ""{
		Pagging.OrderDir = orderDir
	}else{
		Pagging.OrderDir = "DESC"
	}

	
	ResultPaginate := dto.ResultPaginate{
		Offset: Pagging.Offset,
		Limit:Pagging.Limit,
		Order: GenerateOrder(Pagging.OrderBy,Pagging.OrderDir),
	}

	return ResultPaginate, Pagging
	
}

func GenerateOffSate(Page int, PerPage int) int{
	return (Page * PerPage) - PerPage
}



func GenerateOrder(OrderBy string, OrderDir string) string{
	return  fmt.Sprintf("%s %s",OrderBy,OrderDir)
}