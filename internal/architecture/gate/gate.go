// Package gate authenticates Telegram users who trying to use bot
package gate

import (
	"slices"
	"strings"
	"tg-bookmarks-bot/internal/infrastructure/env"
)

// Authenticate checks if user allowed to use bot
func Authenticate(username string) bool {
	if usernames := env.Get("ALLOWED_USERS"); usernames != "" {
		return slices.Contains(strings.Split(usernames, ","), username)
	}

	return false
}
