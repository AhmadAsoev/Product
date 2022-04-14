package repository

import "product/pkg/domain"

func AddProduct(info domain.Product) error {
	return DB.Create(&info).Error
}
