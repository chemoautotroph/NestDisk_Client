package config

import (
	"github.com/spf13/viper"
)

// Config 配置信息
var Config *viper.Viper

func initConfig() {
	v := viper.New()
	v.SetConfigFile(".env")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Config = v
}

func GetConfig() *viper.Viper {
	return Config
}

