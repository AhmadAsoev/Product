package services

import (
	"github.com/google/uuid"
	"net/http"
	"product/pkg/application/repository"
	"product/pkg/domain"
)

func GetById(urlId string) domain.Response {
	if len(urlId) == 0 {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "ID must not be empty",
			Message: nil,
		}
	}
	id, err := uuid.Parse(urlId)
	if err != nil {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: "ID must be uuID format",
		}
	}

	product, err := repository.GetById(id)
	if err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "can get product by id",
			Message: nil,
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: product,
	}
}
