package main

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var viperInit = sync.Once{}

// initViper initializes viper using config.yaml from root dir
func initViper() {
	viperInit.Do(func() {
		viper.SetConfigName(configFileName)
		viper.AddConfigPath(configPath)
		viper.SetConfigType(configType)

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("%s", err))
		}
	})
}

func getKey() string {
	initViper()
	return viper.GetString("key")
}
