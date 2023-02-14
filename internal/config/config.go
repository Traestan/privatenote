package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// path to config file
var configPath = pflag.String("config.path", "./../config.toml", "Path to config file")

// NewConfig returns new instance *viper.Viper with loaded config
func NewConfig() *viper.Viper {
	pflag.Parse()

	config := viper.New()
	config.SetConfigFile(*configPath)
	viper.SetConfigType("toml")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return config
}
