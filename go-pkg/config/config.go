package config

import (
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/joho/godotenv"
)

const (
	PgDatabase = "users"
)

type Config struct {
}

func Init(dev bool) *Config {
	envPath := config.SourcePath + "/github.com/benka-me/users/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}
	ret := &Config{}
	if !dev {
	} else {
	}
	return ret
}
