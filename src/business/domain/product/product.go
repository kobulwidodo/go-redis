package product

import (
	"context"
	"encoding/json"
	"errors"
	"medium-go-redis/src/business/entity"
	"medium-go-redis/src/lib/redis"
	"time"

	"gorm.io/gorm"
)

type Interface interface {
	GetList(ctx context.Context) ([]entity.Product, error)
}

type product struct {
	db    *gorm.DB
	redis redis.Interface
}

func Init(db *gorm.DB, redis redis.Interface) Interface {
	p := &product{
		db:    db,
		redis: redis,
	}

	return p
}

func (p *product) GetList(ctx context.Context) ([]entity.Product, error) {
	res := []entity.Product{}

	// get cache
	cachedProducts, err := p.redis.Get(ctx, entity.ProductListRedisKey)
	switch {
	case errors.Is(err, redis.Nil):
		p.db.Logger.Info(ctx, err.Error())
	case err != nil:
		p.db.Logger.Error(ctx, err.Error())
	}

	if err == nil {
		if err := json.Unmarshal([]byte(cachedProducts), &res); err != nil {
			return res, err
		}
		p.db.Logger.Info(ctx, "successfuly get data from redis")
		return res, nil
	}

	if err := p.db.Find(&res).Error; err != nil {
		return res, err
	}

	// set cache
	productMarshal, err := json.Marshal(res)
	if err != nil {
		p.db.Logger.Error(ctx, err.Error())
	}

	if err := p.redis.SetEx(ctx, entity.ProductListRedisKey, string(productMarshal), time.Hour); err != nil {
		p.db.Logger.Error(ctx, err.Error())
	}

	return res, nil
}
