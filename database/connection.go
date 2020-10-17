package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func ConnectDatabase() error {

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		databaseUrl = os.Getenv("DATABASE_URL")
	}

	once.Do(func() {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			fmt.Println("DATABASE CONNECTION FAILED.")
		}
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(10)
	})
	return err
}

func GetConnection() (*sql.DB, error) {

	return db, err
}
