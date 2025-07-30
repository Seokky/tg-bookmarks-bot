// Package net works with net requests like GET and POST
package net

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"tg-bookmarks-bot/internal/infrastructure/env"
)

var (
	host  = env.Get("TELEGRAM_HOST")
	base  = env.Get("TELEGRAM_BASE")
	token = env.Get("TELEGRAM_API_TOKEN")
)

var client = http.Client{}

// Build url string from host, base, token and entrypoint
func buildURL(entrypoint string) (string, error) {
	if len(host) == 0 || len(base) == 0 || len(token) == 0 {
		log.Fatal("no")
		return "", fmt.Errorf("can't build buildURL with host = %s, base = %s, token = %s", host, base, token)
	}

	result := "https://" + path.Join(host, base+token, entrypoint)

	return result, nil
}
