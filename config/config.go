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

func LoadConfig() config Config, err error {
	return Config{
		DBHost        : "aws-0-ap-southeast-1.pooler.supabase.com"
		DBUserName    : "postgres.kibrvkytbvzpxxssufff" 
		DBUserPassword: "postgressujikom" 
		DBName        : "postgres" 
		DBPort        : "5432" 
		AppSecret     : "qweqwe" 
	}
}
