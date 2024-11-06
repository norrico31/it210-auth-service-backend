package config

import (
	"os"
	"strconv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	DBPort                 string
	GatewayPort            string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "127.0.0.1"),
		Port:                   getEnv("PORT", "8081"),
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", "root"),
		DBAddress:              getEnv("DB_ADDRESS", "postgres"),
		DBPort:                 getEnv("DB_PORT", "5432"),
		GatewayPort:            getEnv("GATEWAY_SERVICE_PORT", "8080"),
		DBName:                 getEnv("DB_NAME", "it210"),
		JWTSecret:              getEnv("JWT_SECRET", "IS-IT_REALL-A_SECRET-?~JWT-NOT_SO-SURE"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		envVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return envVal
	}
	return fallback
}
