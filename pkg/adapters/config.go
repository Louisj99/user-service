package adapters

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	PathToServiceAccountKey string `yaml:"pathToServiceAccountKey"`
	Host                    string `yaml:"HOST"`
	Port                    string `yaml:"PORT"`
	User                    string `yaml:"USER"`
	Password                string `yaml:"PASSWORD"`
	Dbname                  string `yaml:"DBNAME"`
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
