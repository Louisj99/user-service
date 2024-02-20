package adapters

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	PathToServiceAccountKey string `yaml:"pathToServiceAccountKey"`
}

func NewConfig() (*Config, error) {
	var config Config
	err := cleanenv.ReadConfig("config.yaml", &config)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
		return nil, err
	}
	return &config, err
}
