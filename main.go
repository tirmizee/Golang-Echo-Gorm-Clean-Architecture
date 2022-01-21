package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func initializeViper() {

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func init() {
	initializeViper()
}

func main() {

	var server = echo.New()

	server.Logger.Fatal(server.Start(":1323"))
}
