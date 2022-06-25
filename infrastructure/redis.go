package infrastructure

import "github.com/go-redis/redis/v8"

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{Addr: "localhost:6369", Password: "", Db: 0}
}

func (r *RedisConfig) NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.Db,
	})
	return client
}
