package config

import (
	// "fmt"
	"img_tag/utils"
	"os"

	"github.com/spf13/viper"
)

func init() {
	dir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(utils.JoinPath(dir, "config"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
