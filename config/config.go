package config

import (
	"fmt"
	"myapi/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "12345", "myapi", "5432")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal connect database")
	}

	if err := db.AutoMigrate(&models.Register{}, &models.Credentials{}); err != nil {
		log.Fatal("gagal migrate di database", err)
	}
	log.Println("sukses Migrate")

	DB = db
	log.Println("Database Connected")
}
