package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	API_PORT            string
	DB_ADDRESS          string
	DB_USERNAME         string
	DB_PASSWORD         string
	DB_NAME             string
	JWT_SECRET_ACCESS   string
	JWT_SECRET_REFRESH  string
	IMAGEKIT_PRIVKEY    string
	IMAGEKIT_PUBKEY     string
	SUPERADMIN_NAME     string
	SUPERADMIN_EMAIL    string
	SUPERADMIN_PASSWORD string
}

var Env *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	err := viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}

	Env = cfg
}
