package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Job  string
}

type UserRepository interface {
	Create(user User) error
	FindOne(id uint) (User, error)
}

type UserRedisRepository interface {
	Get(key string) (User, error)
	Set(key string, user User, duration int) error
}

type UserUsecase interface {
	Seed() error
	GetByIdWithoutRedis(id uint) (User, error)
	GetByIdWithRedis(id uint) (User, error)
}

type UserIdUriBinding struct {
	Id uint `uri:"id" binding:"required"`
}
