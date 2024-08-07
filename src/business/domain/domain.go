package domain

import (
	"medium-go-redis/src/business/domain/product"
	"medium-go-redis/src/lib/redis"

	"gorm.io/gorm"
)

type Domains struct {
	Product product.Interface
}

func Init(db *gorm.DB, redis redis.Interface) *Domains {
	d := &Domains{
		Product: product.Init(db, redis),
	}

	return d
}
