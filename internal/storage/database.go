package storage

import (
	"log"
	"os"

	"github.com/hazaloolu/openUp_backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}
	log.Println("Database connection established succesfully")

	err = DB.AutoMigrate(&model.User{}, &model.Post{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database schema migrated")

}
