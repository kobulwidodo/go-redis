package main

import (
	"medium-go-redis/src/business/domain"
	"medium-go-redis/src/business/usecase"
	"medium-go-redis/src/handler/rest"
	"medium-go-redis/src/lib/configreader"
	"medium-go-redis/src/lib/redis"
	"medium-go-redis/src/lib/sql"
	"medium-go-redis/src/utils/config"

	_ "medium-go-redis/docs/swagger"
)

// @contact.name   Rakhmad Giffari Nurfadhilah
// @contact.url    https://fadhilmail.tech/
// @contact.email  rakhmadgiffari14@gmail.com

// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
	configFile string = "./etc/cfg/config.json"
)

func main() {
	cfg := config.Init()
	configReader := configreader.Init(configreader.Options{
		ConfigFile: configFile,
	})
	configReader.ReadConfig(&cfg)

	db := sql.Init(cfg.SQL)

	redis := redis.Init(cfg.Redis)

	d := domain.Init(db, redis)

	uc := usecase.Init(d)

	r := rest.Init(cfg.Gin, configReader, uc)

	r.Run()
}
