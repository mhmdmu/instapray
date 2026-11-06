package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	log.Println(os.Getenv("DATABASE_URL"))
	log.Println("InstaPray is alive")
}

func loadEnvFile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading `.env` file")
	}

	log.Println("Loaded `.env` file")
}
