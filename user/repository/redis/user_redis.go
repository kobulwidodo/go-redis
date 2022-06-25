package redis

import (
	"context"
	"encoding/json"
	"go-redis/domain"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserRedisRepository struct {
	redisClient *redis.Client
}

func NewUserRedisRepository(redisClient *redis.Client) domain.UserRedisRepository {
	return &UserRedisRepository{redisClient}
}

func (r *UserRedisRepository) Get(key string) (domain.User, error) {
	userBytes, err := r.redisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{}
	if err := json.Unmarshal(userBytes, &user); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRedisRepository) Set(key string, user domain.User, duration int) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := r.redisClient.Set(context.Background(), key, userBytes, time.Second*time.Duration(duration)).Err(); err != nil {
		return err
	}
	return nil
}
