package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/hazaloolu/openUp_backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	dsn := os.Getenv("DATABASE_URL")

	// dsn := "postgresql://hazaloolu:RPmZ33taUstHSupGy1sxSA1LFPTgZv0x@dpg-cspm8qt2ng1s73d3f2n0-a.oregon-postgres.render.com/openup_db"

	fmt.Print(dsn)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}
	log.Println("Database connection established succesfully")

	err = DB.AutoMigrate(&model.User{}, &model.Post{})
	log.Println("migration completed")

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database schema migrated")

}
