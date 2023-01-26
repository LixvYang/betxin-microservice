package config

import (
	"os"
	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	_ = viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }
}