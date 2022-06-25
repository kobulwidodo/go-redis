package infrastructure

import (
	"fmt"
	"go-redis/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	DbHost string
	DbUser string
	DbPass string
	DbName string
	DbPort string
}

func NewDbConfig() DbConfig {
	return DbConfig{DbHost: "localhost", DbUser: "postgres", DbPass: "password", DbName: "go-redis", DbPort: "5469"}
}

func (c *DbConfig) InitDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.DbHost,
		c.DbUser,
		c.DbPass,
		c.DbName,
		c.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
