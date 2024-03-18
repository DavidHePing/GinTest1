package setup

import (
	"fmt"

	"github.com/spf13/viper"
)

func ViperSetup() {
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Is debug mode!")
	}
}
