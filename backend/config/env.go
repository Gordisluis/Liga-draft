package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port string
	DBUser string
	DBPassword string
	DBAddress string
	DBName string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: Getenv("PUBLIC_HOST", "http://localhost"),
		Port: Getenv("PORT","8080"),
		DBUser: Getenv("DB_USER", "luis"),
		DBPassword: Getenv("DB_PASSWORD", "luis87"),
		DBAddress: fmt.Sprintf("%s:%s", Getenv("DB_HOST", "127.0.0.1"), Getenv("DB_PORT", "3306")),
		DBName: Getenv("DB_NAME", "liga"),
	}
}

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}