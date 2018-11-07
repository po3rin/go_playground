package main

import (
	"fmt"

	"github.com/jmartin82/mconfig"
)

type MysqlConfiguration struct {
	Host     string `env:"MYSQL_HOST"`
	Username string `env:"MYSQL_USERNAME"`
	Password string `env:"MYSQL_PASSWORD"`
	Database string `env:"MYSQL_DATABASE"`
	Port     int    `env:"MYSQL_PORT"`
}

type RedisConfiguration struct {
	Host string `env:"REDIS_HOST"`
	Port int    `env:"REDIS_PORT"`
}

type Configuration struct {
	Port  int `env:"APP_PORT"`
	Mysql MysqlConfiguration
	Redis RedisConfiguration
}

func main() {
	configuration := Configuration{}
	err := mconfig.Parse("conf.yml", &configuration)
	if err != nil {
		panic(err)
	}
	fmt.Println(configuration)
}
