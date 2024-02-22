package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type values struct {
	DatabaseUrl string `envconfig:"DATABASE_URL"`
	APIEndpoint string `envconfig:"API_ENDPOINT" required:"true"`
	LoggerLevel string `envconfig:"LOGGER_LEVEL"`
	SecretJWT   string `envconfig:"SECRET_JWT" required:"true"`
}

var Values values

func LoadFromFile(fPath string) error {
	_ = godotenv.Load(fPath)

	err := envconfig.Process("", &Values)
	if err != nil {
		log.Println("ERROR: envconfig.Process(): ", err.Error())
	}
	return err
}
