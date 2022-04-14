package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Product struct {
	Id          *uuid.UUID `json:"id" gorm:"column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	Price       float64    `json:"price" gorm:"column:price"`
	Description string     `json:"description" gorm:"column:description"`
}

func (Product) TableName() string {
	return "products"
}

func (p Product) Validate() error {

	if len(p.Name) > 15 {
		return errors.New("name must not be more then 15")
	}
	if p.Price == 0 || p.Price < 0 {
		return errors.New("price must not be negative")
	}
	if p.Description == "" || len(p.Description) > 255 {
		return errors.New("description must not empty")
	}
	return nil
}
