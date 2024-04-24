package configs

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Server *Server `mapstructure:"server"`
	Database *Database `mapstructre:"database"`
}	

var Config *config

func InitConfig() {
	if (Config == nil) {
		Config = loadConfigVariables()
	}
}

func loadConfigVariables() (config *config){
	viper.AddConfigPath("environments")
	viper.SetConfigName("environment")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}