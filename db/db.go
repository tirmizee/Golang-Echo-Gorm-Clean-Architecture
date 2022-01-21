package db

import (
	"clean-architect/configs"
	"clean-architect/models"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {

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
		DisableDatetimePrecision:  false,
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

	// migrate table
	db.AutoMigrate(
		&models.User{},
		&models.Role{})

	seedData(db)

	return db
}

func seedData(db *gorm.DB) {

	db.Where("1 = 1").Delete(&models.Role{})
	db.Where("1 = 1").Delete(&models.User{})

	db.Create(&models.User{Username: "tirmizee", Password: "123", Email: "tirmizee@hotmail.com", FirstName: "pratya", LastName: "yeekhaday"})
	db.Create(&models.User{Username: "kiskdifw", Password: "123", Email: "kiskdifw@hotmail.com", FirstName: "poikue", LastName: "poiloipuy"})
	db.Create(&models.Role{Code: "R001", Name: "admin", Desc: "admin"})
	db.Create(&models.Role{Code: "R002", Name: "user", Desc: "user"})
}
