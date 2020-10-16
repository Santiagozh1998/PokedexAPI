package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

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
		databaseUrl = "postgres://bzsdjbvmkgkvxg:a6739ae2dbbdda69fdacb4f04b243f071b0c81b0039ce119ce1193d1f29f0171@ec2-52-204-20-42.compute-1.amazonaws.com:5432/db1l4723kdrhe8"
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
