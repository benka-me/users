package db

import (
	"github.com/benka-me/users/go-pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func Init(config *config.Config, dev bool) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("postgres", config.Postgres(dev))
	if err != nil {
		panic(err)
	}
	return db
}
