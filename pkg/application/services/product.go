package services

import (
	"github.com/google/uuid"
	"net/http"
	"product/pkg/application/repository"
	"product/pkg/domain"
)

func AddProduct(product domain.Product) (response domain.Response) {
	if err := product.Validate(); err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "",
		}
	}
	id := uuid.New()
	product.Id = &id
	if err := repository.AddProduct(product); err != nil {

		return domain.Response{
			Code:    http.StatusInternalServerError,
			Error:   "Can't created product into db",
			Message: "",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Product added",
	}
}
