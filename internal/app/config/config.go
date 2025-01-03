package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"jwtgo/pkg/logging"
)

type Config struct {
	App struct {
		Host  string `yaml:"host" env-required:"true"`
		Port  string `yaml:"port" env-required:"true"`
		Debug bool   `yaml:"debug"`
	} `yaml:"app" env-required:"true"`

	MongoDB struct {
		Url      string `yaml:"url" env-required:"true"`
		Database string `yaml:"database" env-required:"true"`
	} `yaml:"mongodb" env-required:"true"`

	Security struct {
		Salt            string `yaml:"salt" env-required:"true"`
		Secret          string `yaml:"secret" env-required:"true"`
		BcryptCost      int    `yaml:"bcrypt_cost" env-required:"true"`
		AccessLifetime  int    `yaml:"access_lifetime" env-required:"true"`
		RefreshLifetime int    `yaml:"refresh_lifetime" env-required:"true"`
	} `yaml:"security" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logging.Logger) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("configs/config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
