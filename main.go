package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func newGormDB() *gorm.DB {
	return nil

}

func setupMiddlewares(e *echo.Echo) {

}

func setupRoute(e *echo.Echo, db *gorm.DB) {

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
		server = echo.New()
		gormDB = newGormDB()
	)

	setupMiddlewares(server)
	setupRoute(server, gormDB)

	server.Logger.Fatal(server.Start(":1323"))
}
