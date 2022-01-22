package configs

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	dbConfig DBConfigProp
	rdConfig RedisConfigProp
)

func GetDbConfigProp() DBConfigProp {
	return dbConfig
}

func GetRdConfigProp() RedisConfigProp {
	return rdConfig
}

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
	Host  string
	Port  string
	Pass  string
	Index int
}

func initDbConfigProp() {
	dbConfig.Host = viper.GetString("db.host")
	dbConfig.Port = viper.GetString("db.port")
	dbConfig.User = viper.GetString("db.user")
	dbConfig.Pass = viper.GetString("db.pass")
	dbConfig.Dbname = viper.GetString("db.dbname")
	dbConfig.MaxPool = viper.GetInt("db.pool.max")
	dbConfig.IdlePool = viper.GetInt("db.pool.idle")
	dbConfig.MaxLifetime = viper.GetDuration("db.pool.lifetime")
}

func initRdConfigProp() {
	rdConfig.Host = viper.GetString("redis.host")
	rdConfig.Port = viper.GetString("redis.port")
	rdConfig.Pass = viper.GetString("redis.pass")
	rdConfig.Index = viper.GetInt("redis.index")
}

func initializeViper() {

	viper.AddConfigPath(".")      // path
	viper.SetConfigName("config") // filename
	viper.SetConfigType("yaml")   // filetype

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func init() {
	initializeViper()
	initDbConfigProp()
	initRdConfigProp()
}
