package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	return nil

}

func setupRoute(db *gorm.DB) *echo.Echo {

	e := echo.New()

	return e
}

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

	var (
		db     = setupDB()
		server = setupRoute(db)
	)

	server.Logger.Fatal(server.Start(":1323"))
}
