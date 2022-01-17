package main

import (
	"clean-architect/configs"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	dbConfig    configs.DBConfig
	redisConfig configs.RedisConfig
)

func initializeConfig() {

}

func initializeViper() {
	// viper.AddConfigPath("configs")
	// viper.SetConfigName("configs")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func init() {
	fmt.Println("initialize")
	initializeViper()
	initializeConfig()
}

func main() {

}
