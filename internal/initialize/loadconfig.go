package initialize

import (
	"fmt"
	"todolist/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Server Port:", viper.Get("server.port"))
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
