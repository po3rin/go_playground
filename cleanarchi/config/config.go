package config

type Config struct{}

type DB struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}
