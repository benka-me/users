package config

import (
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/joho/godotenv"
	"os"
)

const (
	PgDatabase   = "users"
	PgCollection = "users"
)

type Config struct {
	DbUser string
	DbHost string
	DbPort string
	DbPWD  string
	DbSSL  string
}

func Init(dev bool) *Config {
	envPath := config.SourcePath + "/github.com/benka-me/users/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}
	ret := &Config{
		DbUser: os.Getenv("DB_USER"),
		DbPWD:  os.Getenv("DB_PASSWORD"),
		DbSSL:  os.Getenv("DB_SSLMODE"),
	}
	if !dev {
		ret.DbHost = os.Getenv("DB_HOST_PROD")
		ret.DbPort = os.Getenv("DB_PORT_PROD")
	} else {
		ret.DbHost = os.Getenv("DB_HOST_DEV")
		ret.DbPort = os.Getenv("DB_PORT_DEV")
	}
	return ret
}
