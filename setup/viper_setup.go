package setup

import (
	"GinTest1/domain"
	"fmt"

	"github.com/spf13/viper"
)

func ViperSetup() domain.Config {
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Is debug mode!")
	}

	var config domain.Config
	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	return config
}
