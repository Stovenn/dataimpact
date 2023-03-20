package util

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBName        string        `mapstructure:"DB_NAME"`
	DBUri         string        `mapstructure:"DB_URI"`
	Host          string        `mapstructure:"HOST"`
	Port          string        `mapstructure:"PORT"`
	SymmetricKey  string        `mapstructure:"SYMMETRIC_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
}

func SetupConfig(path string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("could not read from env file")
	}

	_ = viper.Unmarshal(&config)

	return config, nil
}
