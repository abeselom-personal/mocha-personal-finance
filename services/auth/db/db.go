package db

import (
	"database/sql"
	"log"

	"github.com/abeselom-personal/personal-finance/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var SQLDB *sql.DB

func InitDB(cfg config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.ConnString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	SQLDB, err = DB.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm:", err)
	}
}
