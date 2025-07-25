package net

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var host, base, token string
var client = http.Client{}

// Load environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can't load .env file")
	}

	host = os.Getenv("TELEGRAM_HOST")
	base = os.Getenv("TELEGRAM_BASE")
	token = os.Getenv("TELEGRAM_API_TOKEN")
}
