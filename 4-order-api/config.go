package orderapi

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	validationApi "github.com/mrScorpio/dz/3-validation-api"
)

type Config struct {
	Db  DbConfig
	Vld validationApi.VldConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	if godotenv.Load() != nil {
		log.Println("Error loading .env file, use default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
