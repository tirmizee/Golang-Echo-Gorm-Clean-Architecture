package configs

import (
	"time"

	"github.com/spf13/viper"
)

var (
	dbConfig DBConfigProp
	rdConfig RedisConfigProp
)

type DBConfigProp struct {
	Host        string
	Port        string
	User        string
	Pass        string
	Dbname      string
	MaxPool     int
	IdlePool    int
	MaxLifetime time.Duration
}

type RedisConfigProp struct {
	Host string
	Port string
	Pass string
}

func initializeDbConfigProp() {
	dbConfig.Host = viper.GetString("db.host")
	dbConfig.Port = viper.GetString("db.port")
	dbConfig.User = viper.GetString("db.user")
	dbConfig.Pass = viper.GetString("db.pass")
	dbConfig.Dbname = viper.GetString("db.dbname")
	dbConfig.MaxPool = viper.GetInt("db.pool.max")
	dbConfig.IdlePool = viper.GetInt("db.pool.idle")
	dbConfig.MaxLifetime = viper.GetDuration("db.pool.lifetime")
	rdConfig.Host = viper.GetString("")
	rdConfig.Port = viper.GetString("")
	rdConfig.Pass = viper.GetString("")
}

func initializeRdConfigProp() {
	rdConfig.Host = viper.GetString("")
	rdConfig.Port = viper.GetString("")
	rdConfig.Pass = viper.GetString("")
}

func init() {
	initializeDbConfigProp()
	initializeRdConfigProp()
}

func GetDbConfigProp() DBConfigProp {
	return dbConfig
}

func GetRdConfigProp() DBConfigProp {
	return dbConfig
}
