package config

import (
	"fmt"
	"os"
	"strconv"
	"wishlist/util"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	util.ProcessEnv()

	var dsn string
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logrus.Error("missing port!")
		panic(err)
	}

	if os.Getenv("CON_SSL") == "on" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
			host,
			username,
			password,
			dbname,
			port,
		)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			host,
			username,
			password,
			dbname,
			port,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Can't connect postgres database!")
		panic(err)
	}

	MigrateDB(db)

	return db
}
