package config

import (
	"flag"
	"os"

	"app/util"

	"gopkg.in/yaml.v3"
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

type Config struct {
	Server    ServerConfig
	Db        DbConfig
	Migration MigrationConfig
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

	c.Server.Host = util.GetEnv("HOST", c.Server.Host)
	c.Server.Port = util.GetEnv("PORT", c.Server.Port)
	c.Db.Url = util.GetEnv("DB_URL", c.Db.Url)

	return
}
