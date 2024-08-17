package config

import (
	"os"
)

type Config struct {
	DB       DBConfig
	KEYCLOAK KeyCloakConfig
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

type KeyCloakConfig struct {
	Host         string
	Realm        string
	ClientID     string
	ClientSecret string
}

var cfg Config

// Called automatically when this pkg is imported & initialized
func init() {
	//uncomment this part if running locally

	//cwd, err := os.Getwd()
	//if err != nil {
	//	log.Panic(err)
	//}
	//envPath := filepath.Join(cwd, ".env")
	//
	//// Load the .env file
	//err = godotenv.Load(envPath)

	cfg.DB = DBConfig{
		Host:       os.Getenv("DB_HOST"),
		Name:       os.Getenv("DB_NAME"),
		User:       os.Getenv("DB_USER"),
		Pass:       os.Getenv("DB_PASS"),
		Port:       os.Getenv("DB_PORT"),
		MaxOpenCon: os.Getenv("DB_MAX_OPEN_CON"),
		MaxIdleCon: os.Getenv("DB_MAX_IDLE_CON"),
	}

	cfg.KEYCLOAK = KeyCloakConfig{
		Host:         os.Getenv("KEYCLOAK_HOST"),
		Realm:        os.Getenv("KEYCLOAK_REALM"),
		ClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}

}

func New() *Config {
	return &cfg
}
