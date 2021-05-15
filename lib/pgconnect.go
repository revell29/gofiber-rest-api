package pgconnect

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)


func envPGConfig() pgx.ConnConfig {
	err := godotenv.Load(".env")

	if err != nil {
    log.Fatalf("Error loading .env file")
  }


	fmt.Println(os.Getenv("DB_NAME"))
	return pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:   	5432,
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

// dbInitialize -
func Connection()(*pgx.ConnPool, error) {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: envPGConfig(),
		AfterConnect: func(c *pgx.Conn) error {
			return nil
		},
	})

	return pool, err

}