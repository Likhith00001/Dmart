package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	URL string `mapstructure:"url"`
}

type RedisConfig struct {
	URL string `mapstructure:"url"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expiry int    `mapstructure:"expiry"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
}

var (
	cfg  Config
	once sync.Once
)

func Load() Config {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/app") // for Docker
		viper.AutomaticEnv()

		// Bind environment variables explicitly
		viper.BindEnv("SERVER_PORT", "SERVER_PORT")
		viper.BindEnv("SERVER_ENV", "SERVER_ENV")
		viper.BindEnv("DATABASE_URL", "DATABASE_URL")
		viper.BindEnv("REDIS_URL", "REDIS_URL")
		viper.BindEnv("JWT_SECRET", "JWT_SECRET")
		viper.BindEnv("JWT_EXPIRY", "JWT_EXPIRY")

		if err := viper.ReadInConfig(); err == nil {
			log.Printf("Loaded config from: %s", viper.ConfigFileUsed())
		} else {
			log.Println("No .env file found, using environment variables")
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("Failed to unmarshal config: %v", err)
		}

		// Set defaults
		if cfg.Server.Port == "" {
			cfg.Server.Port = "8081"
		}
		if cfg.Server.Env == "" {
			cfg.Server.Env = "development"
		}
		if cfg.Database.URL == "" {
			cfg.Database.URL = "postgres://postgres:postgres@postgres:5432/user_db?sslmode=disable"
		}
	})

	return cfg
}
