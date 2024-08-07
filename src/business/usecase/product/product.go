package product

import (
	"context"
	productDom "medium-go-redis/src/business/domain/product"
	"medium-go-redis/src/business/entity"
)

type Interface interface {
	GetList(ctx context.Context) ([]entity.Product, error)
}

type product struct {
	product productDom.Interface
}

func Init(pd productDom.Interface) Interface {
	p := &product{
		product: pd,
	}

	return p
}

func (p *product) GetList(ctx context.Context) ([]entity.Product, error) {
	products, err := p.product.GetList(ctx)
	if err != nil {
		return products, err
	}

	return products, nil
}
