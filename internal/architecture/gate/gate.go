// Package gate authenticates Telegram users who trying to use bot
package gate

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"slices"
	"strings"
)

// Authenticate checks if user allowed to use bot
func Authenticate(username string) bool {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can't load .env file")
	}

	usernames := os.Getenv("ALLOWED_USERS")

	if usernames == "" {
		return false
	}

	return slices.Contains(strings.Split(usernames, ","), username)
}
