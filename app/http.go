package main

import (
	"go-redis/infrastructure"
	_userHandler "go-redis/user/delivery/http"
	_userRepository "go-redis/user/repository/postgresql"
	_userRedisRepository "go-redis/user/repository/redis"
	_userUsecase "go-redis/user/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbDriver := infrastructure.NewDbConfig()
	db, err := dbDriver.InitDb()
	if err != nil {
		log.Fatal("failed to connect database")
		panic(err)
	}

	redisDriver := infrastructure.NewRedisConfig()
	redisClient := redisDriver.NewRedisClient()

	r := gin.Default()
	api := r.Group("/api")

	userRepository := _userRepository.NewUserRepository(db)
	userRedisRepository := _userRedisRepository.NewUserRedisRepository(redisClient)
	userUsecase := _userUsecase.NewUserUsecase(userRepository, userRedisRepository)
	_userHandler.NewUserHandler(api, userUsecase)

	r.Run()
}
