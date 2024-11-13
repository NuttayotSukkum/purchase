package db_connector

import (
	"fmt"
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDb(username, password, host, database, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		database,
		password,
	)
	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panicf("Error Database connect: %s", err)
	}

	if err := db.AutoMigrate(&entities.Product{}, &entities.Payment{}); err != nil {
		log.Errorf("Unable to migrate database")
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Panicf("Failed to create connection pools: %s", err)
	}
	log.Printf("Database is initialized: %v", sqlDb.Stats())
	return db
}
