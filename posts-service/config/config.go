package config

import (
	"github.com/joho/godotenv"
	"os"
)

type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type DbConfig struct {
	Url string `mapstructure:"DB_URL"`
}

type MigrationConfig struct {
	Url string `mapstructure:"MIGRATION_URL"`
}

type AuthConfig struct {
	JwtSigningKey string `mapstructure:"JWT_SIGNING_KEY"`
}

type Config struct {
	Server    ServerConfig
	Db        DbConfig
	Migration MigrationConfig
	Auth      AuthConfig
}

func LoadConfig() (c Config, err error) {
	godotenv.Load(".env")

	c.Server.Host = os.Getenv("SERVER_HOST")
	c.Server.Port = os.Getenv("SERVER_PORT")

	c.Db.Url = os.Getenv("DB_URL")
	c.Migration.Url = os.Getenv("MIGRATION_URL")
	c.Auth.JwtSigningKey = os.Getenv("JWT_SIGNING_KEY")

	return
}
