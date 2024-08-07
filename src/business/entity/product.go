package entity

import (
	"gorm.io/gorm"
)

const (
	ProductListRedisKey = "product_list"
)

type Product struct {
	gorm.Model
	Name        string
	ImgUrl      string
	Description string
	Price       uint64
}
