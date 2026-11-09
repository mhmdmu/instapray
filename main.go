package main

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()

	pool := connectDB()
	defer pool.Close()

	log.Println("InstaPray is alive")
}

func loadEnvFile() {
	err := godotenv.Load()
	if errors.Is(err, fs.ErrNotExist) {
		log.Println("No .env file, using existing environment variables")
		return
	}
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Loaded .env file")
}

func connectDB() *pgxpool.Pool {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Fatalf("Invalid DATABASE_URL: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		log.Fatalf("Cannot reach database: %v", err)
	}

	return pool
}
