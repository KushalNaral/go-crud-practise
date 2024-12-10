package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	DB_TYPE string
	DB_URL  string
}

func LoadConfig() *DBConfig {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	return &DBConfig{
		DB_TYPE: os.Getenv("DB_TYPE"),
		DB_URL:  os.Getenv("DB_URL"),
	}
}

func ConnectDatabase() (*sql.DB, error) {

	db, err := sql.Open(LoadConfig().DB_TYPE, LoadConfig().DB_URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database : %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping to the database : %v", err)
	}

	return db, nil
}
