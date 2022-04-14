package repository

import "product/pkg/domain"

func GetAllProducts() (products []domain.Product, err error) {
	res := DB.Find(&products)
	return products, res.Error
}
