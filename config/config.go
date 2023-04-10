package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func InitConfiguration() Configuration {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return Configuration{
		Host:     viper.GetString("Host"),
		Port:     viper.GetString("Port"),
		Username: viper.GetString("Username"),
		Password: viper.GetString("Password"),
		Name:     viper.GetString("Name"),
	}
}
