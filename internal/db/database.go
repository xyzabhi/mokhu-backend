package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() *sql.DB {
	// Use default connection string if DATABASE_URL is not set
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@127.0.0.1:5432/goapi?sslmode=disable"
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Database unreachable: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
	return db
}
