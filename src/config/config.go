package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host       string
	Name       string
	User       string
	Pass       string
	Port       string
	MaxOpenCon string
	MaxIdleCon string
}

var cfg Config

// Called automatically when this pkg is imported & initialized
func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	cfg.DB = DBConfig{
		Host:       os.Getenv("DB_HOST"),
		Name:       os.Getenv("DB_NAME"),
		User:       os.Getenv("DB_USER"),
		Pass:       os.Getenv("DB_PASS"),
		Port:       os.Getenv("DB_PORT"),
		MaxOpenCon: os.Getenv("DB_MAX_OPEN_CON"),
		MaxIdleCon: os.Getenv("DB_MAX_IDLE_CON"),
	}

}

func New() *Config {
	return &cfg
}
