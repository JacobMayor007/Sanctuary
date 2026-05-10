package database

import (
	"log"
	"os"
	"sanctuary_server/model"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbcn := os.Getenv("DATABASE_URL")

	// Try to connect first
	db, err := gorm.Open(postgres.Open(dbcn), &gorm.Config{})
	if err != nil {
		// If "sanctuary" doesn't exist, create it
		if strings.Contains(err.Error(), "database") && strings.Contains(err.Error(), "does not exist") {
			log.Println("Database does not exist, creating...")
			createDatabase(dbcn)

			// Retry connection after creating
			db, err = gorm.Open(postgres.Open(dbcn), &gorm.Config{})
			if err != nil {
				log.Fatal("Failed to connect to database:", err)
			}
		} else {
			log.Fatal("Failed to connect to database:", err)
		}
	}

	log.Println(`
	 =========================================
	║                                         ║
	║	 Sanctuary Database Connected!    ║
	║                                         ║
	║             v1.0.0                      ║
	║                                         ║
	║                                         ║
	║                                         ║
	 =========================================
	`)

	log.Println("Running migrations...")
	db.AutoMigrate(&model.User{})
	log.Println("Migrations completed successfully!")
	DB = db
}

func createDatabase(dbcn string) {
	// Connect to the default "postgres" database to run CREATE DATABASE
	// Replace the database name in the DSN with "postgres"
	adminDSN := replaceDatabaseName(dbcn, "postgres")

	db, err := gorm.Open(postgres.Open(adminDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to postgres database:", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	_, err = sqlDB.Exec("CREATE DATABASE sanctuary;")
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	log.Println("Database 'sanctuary' created successfully!")
}

func replaceDatabaseName(dsn, newDBName string) string {
	// DSN format: postgres://user:password@host:port/dbname?options
	parts := strings.Split(dsn, "/")
	// Replace the last part (dbname?options) with newDBName
	parts[len(parts)-1] = newDBName
	return strings.Join(parts, "/")
}
