package usecase

import (
	"medium-go-redis/src/business/domain"
	"medium-go-redis/src/business/usecase/product"
)

type Usecase struct {
	Product product.Interface
}

func Init(d *domain.Domains) *Usecase {
	uc := &Usecase{
		Product: product.Init(d.Product),
	}

	return uc
}
