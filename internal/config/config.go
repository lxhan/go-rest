package config

import (
	"go-rest/pkg/log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Dev  *bool  `env:"DEV"`
	Host string `env:"HOST"`
	Port string `env:"PORT" env-default:"8888"`
}

var conf *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := log.GetLogger()
		logger.Info("reading config")
		conf = &Config{}
		if err := cleanenv.ReadEnv(conf); err != nil {
			logger.Fatal(err)
		}
	})
	return conf
}
