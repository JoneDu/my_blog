package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbDriver  string
	DbDsn     string
	JWTSecret string
}

func LoadConfig() Config {
	// 加载文件
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		DbDriver:  getEnv("DB_DRIVER", "mysql"),
		DbDsn:     getEnv("DB_DSN", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if v := os.Getenv(key); v == "" {
		return defaultVal
	} else {
		return v
	}
}
