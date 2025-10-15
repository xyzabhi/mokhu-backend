package main

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/xyzabhi/mokhu-backend/internal/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	conn := db.Connect()
	defer conn.Close()

	log.Println("🚀 App is ready and DB connection is active")
}
