package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"PGHOST"`
	DBUserName     string `mapstructure:"PGUSER"`
	DBUserPassword string `mapstructure:"PGPASSWORD"`
	DBName         string `mapstructure:"PGNAME"`
	DBPort         string `mapstructure:"PGPORT"`
	AppSecret      string `mapstructure:"APP_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
