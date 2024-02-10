package adapters

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type config struct {
	PathToServiceAccountKey string `yaml:"pathToServiceAccountKey"`
}

func NewConfig() (*config, error) {
	var config config
	err := cleanenv.ReadConfig("config.yaml", &config)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
		return nil, err
	}
	return &config, err
}
