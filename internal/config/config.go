package config

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	App struct {
		Environment string
		Version     string
		Addr        string
		Port        string
	}

	Storage struct {
		Name     string
		Username string
		Password string
		Addr     string
		SSL      string
	}

	Auth struct {
		AuthKey              string
		TokenAuth            *jwtauth.JWTAuth
		GoogleClientID       string
		GoogleClientSecret   string
		GithubClientID       string
		GithubClientSecret   string
		FacebookClientID     string
		FacebookClientSecret string
	}
}

var (
	config *Config
	once   sync.Once
)

func New() {
	once.Do(func() {
		var cfg Config

		if os.Getenv("APP_ENV") == "dev" {
			err := godotenv.Load()
			if err != nil {
				config = &cfg
				return
			}
		}

		cfg.App.Environment = getEnv("APP_ENV", "dev")
		cfg.App.Version = getEnv("APP_VERSION", "0")
		cfg.App.Addr = getEnv("APP_ADDR", "localhost")
		cfg.App.Port = getEnv("APP_PORT", "8080")
		cfg.Storage.Name = getEnv("APP_DB_NAME", "goshop")
		cfg.Storage.Username = getEnv("APP_DB_USERNAME", "user")
		cfg.Storage.Password = getEnv("APP_DB_PASSWORD", "pass")
		cfg.Storage.Addr = getEnv("APP_DB_ADDR", "localhost:5432")
		cfg.Storage.SSL = getEnv("APP_DB_SSL", "disable")
		cfg.Auth.AuthKey = getEnv("APP_AUTH_KEY", "default")
		cfg.Auth.GoogleClientID = getEnv("APP_GOOGLE_CLIENT_ID", "default")
		cfg.Auth.GoogleClientSecret = getEnv("APP_GOOGLE_CLIENT_SECRET", "default")
		cfg.Auth.GithubClientID = getEnv("APP_GITHUB_CLIENT_ID", "default")
		cfg.Auth.GithubClientSecret = getEnv("APP_GITHUB_CLIENT_SECRET", "default")
		cfg.Auth.FacebookClientID = getEnv("APP_FACEBOOK_CLIENT_ID", "default")
		cfg.Auth.FacebookClientSecret = getEnv("APP_FACEBOOK_CLIENT_SECRET", "default")

		config = &cfg
	})
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf(
			"%s environment variable is not defined, so default value %s is used",
			key,
			defaultValue,
		)
		return defaultValue
	}
	return value
}

func Get() *Config {
	return config
}

func (c *Config) GetBaseURL() string {
	protocol := "https://"
	basePath := fmt.Sprintf("%s/api/v%s", c.App.Addr, c.App.Version)

	if c.App.Environment == "dev" {
		protocol = "http://"
		basePath = fmt.Sprintf("%s:%s/api/v%s", c.App.Addr, c.App.Port, c.App.Version)
	}

	return fmt.Sprintf("%s%s", protocol, basePath)
}
