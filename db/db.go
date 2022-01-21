package db

import (
	"clean-architect/configs"
	"clean-architect/models"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GormDB() *gorm.DB {
	return db
}

func initDB() {

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
}

func initMigration() {
	db.AutoMigrate(
		&models.User{},
		&models.Role{})
}

func init() {
	initDB()
	initMigration()
}
