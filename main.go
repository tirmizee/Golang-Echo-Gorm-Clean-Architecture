package main

import (
	"clean-architect/configs"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newGormDB() *gorm.DB {

	// dns config
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.GetDbConfigProp().User,
		configs.GetDbConfigProp().Pass,
		configs.GetDbConfigProp().Host,
		configs.GetDbConfigProp().Port,
		configs.GetDbConfigProp().Dbname)

	// mysql config
	mysqlConfig := mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	// gorm config
	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		DryRun:                 false,
		PrepareStmt:            true,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		panic(err)
	}

	// connection pool config
	if sqlDB, err := db.DB(); err == nil {

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

	}

	return db

}

func setupMiddleware(e *echo.Echo) {

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

	setupMiddleware(server)
	setupRoute(server, gormDB)

	server.Logger.Fatal(server.Start(":1323"))
}
