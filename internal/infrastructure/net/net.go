package net

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path"
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

// Build url string from host, base, token and entrypoint
func buildURL(entrypoint string) (string, error) {
	if len(host) == 0 || len(base) == 0 || len(token) == 0 {
		return "", fmt.Errorf("can't build buildURL with host = %s, base = %s, token = %s", host, base, token)
	}

	result := "https://" + path.Join(host, base+token, entrypoint)

	return result, nil
}
