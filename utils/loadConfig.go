package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
}

func LoadConfig() Config {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	var config Config
	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	return config
}
