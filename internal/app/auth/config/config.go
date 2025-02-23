package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"jwtgo/pkg/logging"
)

type Config struct {
	App struct {
		Host string `env:"AUTH_SERVICE_HOST,required"`
		Port string `env:"AUTH_SERVICE_PORT,required"`
	} `env-required:"true"`

	Service struct {
		User struct {
			Host      string `env:"USER_SERVICE_HOST,required"`
			Port      string `env:"USER_SERVICE_PORT,required"`
			Container string `env:"USER_SERVICE_CONTAINER,required"`
		}
	}

	Security struct {
		Salt            string `env:"SECURITY_SALT,required"`
		Secret          string `env:"SECURITY_SECRET,required"`
		BcryptCost      int    `env:"SECURITY_BCRYPT_COST,required"`
		AccessLifetime  int    `env:"SECURITY_ACCESS_LIFETIME,required"`
		RefreshLifetime int    `env:"SECURITY_REFRESH_LIFETIME,required"`
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
