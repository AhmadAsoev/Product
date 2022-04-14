package repository

import (
	"github.com/google/uuid"
	"product/pkg/domain"
)

func GetById(id uuid.UUID) (product domain.Product, err error) {
	res := DB.First(&product, id)
	return product, res.Error
}
