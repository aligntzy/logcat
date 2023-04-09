package config

import "github.com/spf13/viper"

type config struct {
	Path []string
}

var C config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../configs/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
