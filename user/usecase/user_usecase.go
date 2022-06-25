package usecase

import (
	"fmt"
	"go-redis/domain"
)

const (
	basePrefix    = "user:"
	cacheDuration = 3600
)

type UserUsecase struct {
	userRepository      domain.UserRepository
	userRedisRepository domain.UserRedisRepository
}

func NewUserUsecase(userRepository domain.UserRepository, userRedisReposiotry domain.UserRedisRepository) domain.UserUsecase {
	return &UserUsecase{userRepository, userRedisReposiotry}
}

func (u *UserUsecase) Seed() error {
	for i := 0; i < 100; i++ {
		user := domain.User{
			Name: fmt.Sprintf("User %d", i),
			Job:  "Student",
		}
		if err := u.userRepository.Create(user); err != nil {
			return err
		}
	}
	return nil
}

func (u *UserUsecase) GetByIdWithoutRedis(id uint) (domain.User, error) {
	user, err := u.userRepository.FindOne(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserUsecase) GetByIdWithRedis(id uint) (domain.User, error) {
	cachedUser, err := u.userRedisRepository.Get(u.GenerateKey(id))
	if err != nil {
		fmt.Println(err.Error())
	}
	if cachedUser.ID != 0 {
		return cachedUser, nil
	}
	user, err := u.userRepository.FindOne(id)
	if err != nil {
		return user, err
	}
	if err := u.userRedisRepository.Set(u.GenerateKey(id), user, cacheDuration); err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserUsecase) GenerateKey(id uint) string {
	return fmt.Sprintf("%s:%d", basePrefix, id)
}
