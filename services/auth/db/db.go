package db

import (
	"database/sql"
	"log"

	"github.com/abeselom-personal/personal-finance/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg config.Config) {
	var err error
	DB, err = sql.Open("postgres", cfg.ConnString)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("failed to ping database:", err)
	}
}
