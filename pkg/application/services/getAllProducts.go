package services

import (
	"net/http"
	"product/pkg/application/repository"
	"product/pkg/domain"
)

func GetAllProducts() domain.Response {

	products, err := repository.GetAllProducts()
	if err != nil {
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "Can't get product info into db",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: products,
	}

}
