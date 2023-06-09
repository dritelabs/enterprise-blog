package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	MONGOUri          string `mapstructure:"MONGODB_URI"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
