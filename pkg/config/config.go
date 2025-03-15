package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/npvu1510/en-vocab-server/pkg/utils"
)

type Config struct {
	App struct {
		Port int `env:"APP_PORT" envDefault:"3000"`
	}

	Postgres struct {
		Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
		Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
		User     string `env:"POSTGRES_USER" envDefault:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:""`
		DbName   string `env:"POSTGRES_DBNAME" envDefault:""`
	}
}

var (
	conf *Config
	once sync.Once
	lock sync.Mutex
)

func init() {
	once.Do(func() {
		if !utils.IsLocal() {
			return
		}

		err := godotenv.Load(".env")
		utils.PanicError(err)
	})
}

func MustLoad() *Config {
	lock.Lock()
	defer lock.Unlock()

	if conf != nil {
		return conf
	}

	var config Config
	err := env.Parse(&config)
	utils.PanicError(err)

	conf = &config
	return conf
}
