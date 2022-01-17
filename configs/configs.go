package configs

import "time"

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
