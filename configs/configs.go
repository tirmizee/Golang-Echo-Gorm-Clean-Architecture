package configs

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	DbConfig DBConfig
	RdConfig RedisConfig
)

type DBConfig struct {
	Host        string
	Port        string
	User        string
	Pass        string
	MaxPool     int
	IdlePool    int
	MaxLifetime time.Duration
}

type RedisConfig struct {
	Host string
	Port string
}

func initializeConfigProp() {
	DbConfig.Host = viper.GetString("db.host")
	DbConfig.Port = viper.GetString("db.port")
	DbConfig.User = viper.GetString("db.user")
	DbConfig.Pass = viper.GetString("db.pass")
	DbConfig.MaxPool = viper.GetInt("db.pool.max")
	DbConfig.IdlePool = viper.GetInt("db.pool.idle")
	DbConfig.MaxLifetime = viper.GetDuration("db.pool.lifetime")
	RdConfig.Host = viper.GetString("")
	RdConfig.Port = viper.GetString("")
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
	initializeConfigProp()
}
