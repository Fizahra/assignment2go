package config

import (
	"assignment2go/entity"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	//manggil file env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file")
	}
	//inisialisasi variabel-variabel env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//koneksi ke databasenya
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	//bikin tabel database pake struct yang ada di entity
	db.AutoMigrate(entity.Order{}, entity.Item{}, entity.User{})
	return db
}

// kalo gagal mutus koneksi sama database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection with database")
	}
	dbSQL.Close()
}
