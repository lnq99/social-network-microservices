package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type ServiceConfig struct {
	ProfilesAddr string `yaml:"profiles_addr"`
	PostsAddr    string `yaml:"posts_addr"`
}

type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type AuthConfig struct {
	JwtSigningKey string `mapstructure:"JWT_SIGNING_KEY"`
}

type Config struct {
	Service ServiceConfig
	Server  ServerConfig
	Auth    AuthConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadConfig() (c *Config, err error) {
	filename := flag.String("configFile", "config.yaml", "Config file (default: config.yaml)")

	data, err := os.ReadFile(*filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return
	}

	c.Server.Host = getEnv("HOST", c.Server.Host)
	c.Server.Port = getEnv("PORT", c.Server.Port)
	c.Auth.JwtSigningKey = getEnv("JWT_SIGNING_KEY", c.Auth.JwtSigningKey)

	return
}
