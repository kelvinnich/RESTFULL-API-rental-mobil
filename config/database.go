package config

import (
	"fmt"
	"log"
	
	"restfull-api-rental-mobil/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB{

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("failed to load env file", errEnv)
	}

	// host := os.Getenv("HOST")
	// port := os.Getenv("PORT")
	// user := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	// db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("root:Kutilonta123@@tcp(localhost:3306)/restfull_api?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	
	db.AutoMigrate(&models.User{}, &models.Menu{}, &models.Transaction{})
	return db
}

func CloseDB(db *gorm.DB){
	dbSQL, err := db.DB()

	if err != nil {
		log.Fatal("Failed to close database!")
	}

	dbSQL.Close()
}