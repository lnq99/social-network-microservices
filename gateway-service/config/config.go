package config

import (
	"github.com/joho/godotenv"
	"os"
)

type ServiceConfig struct {
	ProfilesAddr string `yaml:"profiles_addr"`
	PostsAddr    string `yaml:"posts_addr"`
	StatsAddr    string `yaml:"stats_addr"`
	QueueAddr    string `yaml:"queue_addr"`
}

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
	Service   ServiceConfig
	Server    ServerConfig
	Db        DbConfig
	Migration MigrationConfig
	Auth      AuthConfig
}

func LoadConfig() (c Config, err error) {
	godotenv.Load(".env")

	c.Server.Host = os.Getenv("SERVER_HOST")
	c.Server.Port = os.Getenv("SERVER_PORT")

	c.Service.ProfilesAddr = os.Getenv("SERVICE_PROFILES_ADDR")
	c.Service.PostsAddr = os.Getenv("SERVICE_POSTS_ADDR")
	c.Service.StatsAddr = os.Getenv("SERVICE_STATS_ADDR")
	c.Service.QueueAddr = os.Getenv("SERVICE_QUEUE_ADDR")

	c.Db.Url = os.Getenv("DB_URL")
	c.Migration.Url = os.Getenv("MIGRATION_URL")
	c.Auth.JwtSigningKey = os.Getenv("JWT_SIGNING_KEY")

	return
}

//func getEnv(key, fallback string) string {
//	if value, ok := os.LookupEnv(key); ok {
//		return value
//	}
//	return fallback
//}

//func LoadConfig() (c *Config, err error) {
//	viper.AddConfigPath(path)
//	viper.SetConfigFile(".env")
//	viper.AutomaticEnv()
//
//	err = viper.ReadInConfig()
//	if err != nil {
//	return
//	}
//	err = viper.Unmarshal(&config)
//	return

//	filename := flag.String("configFile", "config.yaml", "Config file (default: config.yaml)")
//
//	data, err := os.ReadFile(*filename)
//	if err != nil {
//		return
//	}
//	err = yaml.Unmarshal(data, &c)
//	if err != nil {
//		return
//	}
//
//	c.Server.Host = getEnv("HOST", c.Server.Host)
//	c.Server.Port = getEnv("PORT", c.Server.Port)
//	c.Auth.JwtSigningKey = getEnv("JWT_SIGNING_KEY", c.Auth.JwtSigningKey)
//
//	return
//}
