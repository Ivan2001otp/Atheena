package config

import (
	"log"
	"os"
	"sync"
	"github.com/joho/godotenv"
)

// Add even more configs here if required and update changes.
type Config struct {
	AppEnvironment string
	Port string
	DB_Port string
	JWT_Secret string
}

var (
	instance *Config
	once sync.Once
)

func getEnv(key, fallback string)  string {
	if value, exists := os.LookupEnv(key); exists {
		return value;
	}

	return fallback;
}

func LoadEnvConfig() *Config{
	once.Do(func() {

		if  err := godotenv.Load(".env"); err != nil {
			log.Println(err.Error())
			log.Fatal("⚠️ Failed to read .env file. Something wrong with path or **");
		}

		instance = &Config{
			AppEnvironment:getEnv("APP_ENV_TYPE", "dev"),
			Port: getEnv("PORT", "8080"),
			DB_Port: getEnv("DB_PORT", "27017"),
			JWT_Secret: getEnv("SECRET_KEY","simple-121-s23dif3Af31"),
		}
	});

	return instance;
} 

