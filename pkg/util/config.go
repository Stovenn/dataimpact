package util

// import (
// 	"fmt"
// 	"time"

// 	"github.com/spf13/viper"
// )

// // Config stores all the configuration of Gotodo
// type Config struct {
// 	DBDriver      string        `mapstructure:"DB_DRIVER"`
// 	DBUrl         string        `mapstructure:"DB_URL"`
// 	Host          string        `mapstructure:"HOST"`
// 	Port          string        `mapstructure:"PORT"`
// 	SymmetricKey  string        `mapstructure:"SYMMETRIC_KEY"`
// 	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
// }

// // SetupConfig find a local config file and map the keys to a Config struct
// func SetupConfig(path string) (Config, error) {
// 	var config Config
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		return Config{}, fmt.Errorf("could not read from env file")
// 	}

// 	err = viper.Unmarshal(&config)

// 	return config, nil
// }
