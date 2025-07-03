package main

import (
	"log"

	"github.com/abeselom-personal/personal-finance/config"
	"github.com/abeselom-personal/personal-finance/db"
	"github.com/abeselom-personal/personal-finance/models"
	"github.com/abeselom-personal/personal-finance/routes"
	"github.com/gin-gonic/gin"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// @title Personal Finance Auth API
// @version 1.0
// @description This is the Personal Finance Auth Microservice API.
// @host localhost
// @BasePath /auth
func main() {
	config := config.LoadConfig()
	db.InitDB(config)

	driver, err := postgres.WithInstance(db.SQLDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("error initializing migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("error creating migration instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("error running migrations: %v", err)
	}
	if err == nil {
		log.Println("migrations ran successfully")
	}
	if err := db.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Token{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()
	routes.RegisterRoutes(router)

	err = router.Run()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
