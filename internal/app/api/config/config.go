package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"jwtgo/pkg/logging"
)

type Config struct {
	Service struct {
		Api struct {
			Host  string `env:"API_GATEWAY_HOST,required"`
			Port  string `env:"API_GATEWAY_PORT,required"`
			Debug bool   `env:"API_GATEWAY_DEBUG"`
		} `env-required:"true"`
		Auth struct {
			Host      string `env:"AUTH_SERVICE_HOST,required"`
			Port      string `env:"AUTH_SERVICE_PORT,required"`
			Container string `env:"AUTH_SERVICE_CONTAINER,required"`
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
