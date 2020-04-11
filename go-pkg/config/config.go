package config

import (
	"fmt"
	"github.com/benka-me/laruche/go-pkg/config"
	"github.com/joho/godotenv"
	"os"
)

const (
	PgDatabase = "users"
)

type Config struct {
	PgPort string
	PgHost string
	pgUser string
	pgPwd  string
	pgSSL  string
}

func Init(dev bool) *Config {
	envPath := config.SourcePath + "/github.com/benka-me/users/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}
	ret := &Config{
		PgPort: os.Getenv("POSTGRES_PORT"),
		pgUser: os.Getenv("POSTGRES_USER"),
		pgPwd:  os.Getenv("POSTGRES_PASSWORD"),
		pgSSL:  os.Getenv("POSTGRES_SSLMODE"),
	}
	if !dev {
		ret.PgHost = os.Getenv("POSTGRES_HOST_DEV")
	} else {
		ret.PgHost = os.Getenv("POSTGRES_HOST_PROD")
	}
	return ret
}

func (c Config) Postgres(dev bool) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.PgHost, c.PgPort, c.pgUser, PgDatabase, c.pgPwd, c.pgSSL)
}
