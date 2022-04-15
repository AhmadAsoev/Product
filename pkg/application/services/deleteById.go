package services

import (
	"github.com/google/uuid"
	"net/http"
	"product/pkg/application/repository"
	"product/pkg/domain"
)

func DeleteById(urlId string) domain.Response {
	if len(urlId) == 0 {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "Id must not be empty",
			Message: nil,
		}
	}
	id, err := uuid.Parse(urlId)
	if err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "Can not parse into uuid",
			Message: nil,
		}
	}
	err = repository.DeleteById(id)
	if err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "can't delete product by id",
			Message: nil,
		}
	}
	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "successful delete",
	}
}
