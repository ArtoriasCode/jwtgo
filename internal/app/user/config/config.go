package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"jwtgo/pkg/logging"
)

type Config struct {
	App struct {
		Host string `env:"USER_SERVICE_HOST,required"`
		Port string `env:"USER_SERVICE_PORT,required"`
	} `env-required:"true"`

	MongoDB struct {
		Uri      string `env:"MONGODB_URI,required"`
		Host     string `env:"MONGODB_HOST,required"`
		Port     int    `env:"MONGODB_PORT,required"`
		User     string `env:"MONGODB_USER,required"`
		Password string `env:"MONGODB_PASSWORD,required"`
		Database string `env:"MONGODB_DATABASE,required"`
	} `env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logging.Logger) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(".env", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
