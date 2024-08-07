package config

import (
	"medium-go-redis/src/lib/redis"
	"medium-go-redis/src/lib/sql"
	"time"
)

type Application struct {
	Meta  ApplicationMeta
	Gin   GinConfig
	SQL   sql.Config
	Redis redis.Config
}

type ApplicationMeta struct {
	Title       string
	Description string
	Host        string
	BasePath    string
	Version     string
}

type GinConfig struct {
	Port            string
	Mode            string
	Timeout         time.Duration
	ShutdownTimeout time.Duration
	CORS            CORSConfig
	Meta            ApplicationMeta
}

type CORSConfig struct {
	Mode string
}

func Init() Application {
	return Application{}
}
