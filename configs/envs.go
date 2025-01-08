package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
	CreateApiGithub        string
	RawUrl                 string
	PutApiGithub           string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "5000"),
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", "password"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "http://localhost"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "cms"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		CreateApiGithub:        getEnv("CREATE_API_GITHUB", "https://api.github.com/user"),
		RawUrl:                 getEnv("RAW_URL", "https://raw.githubusercontent.com"),
		PutApiGithub:           getEnv("PUT_API_GITHUB", "https://api.github.com"),
	}


	// return Config{
	// 	PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
	// 	Port:                   getEnv("PORT", "5000"),
	// 	DBUser:                 getEnv("DB_USER", "postgres"),
	// 	DBPassword:             getEnv("DB_PASSWORD", "password"),
	// 	DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "http://localhost"), getEnv("DB_PORT", "3306")),
	// 	DBName:                 getEnv("DB_NAME", "cms"),
	// 	JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
	// 	JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	// 	CreateApiGithub:        getEnv("CREATE_API_GITHUB", "https://api.github.com/user/repos"),
	// 	RawUrl:                 getEnv("RAW_URL", "https://raw.githubusercontent.com"),
	// 	PutApiGithub:           getEnv("PUT_API_GITHUB", "https://api.github.com/repos/%v/%v/contents/%v"),
	// }
}

func getEnv(key, f string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return f
}

func getEnvAsInt(key string, f int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return f
		}
		return i
	}
	return f
}
