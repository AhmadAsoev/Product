package repository

import (
	"github.com/google/uuid"
	"product/pkg/domain"
)

func DeleteById(id uuid.UUID) (err error) {
	var product domain.Product
	res := DB.Delete(&product, id)
	return res.Error
}
