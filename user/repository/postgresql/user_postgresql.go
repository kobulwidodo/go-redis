package postgresql

import (
	"go-redis/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user domain.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindOne(id uint) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
