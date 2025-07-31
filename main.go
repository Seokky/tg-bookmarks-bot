// Telegram bot that stores links passed to it and allow to pick random
package main

import (
	"tg-bookmarks-bot/internal/architecture/handler"
	"tg-bookmarks-bot/internal/architecture/poller"
	"tg-bookmarks-bot/internal/domain/entities"
	"tg-bookmarks-bot/internal/infrastructure/env"
	"tg-bookmarks-bot/internal/infrastructure/net"
)

func main() {
	client := net.NewClient(
		env.Get("TELEGRAM_HOST"),
		env.Get("TELEGRAM_BASE"),
		env.Get("TELEGRAM_API_TOKEN"),
		"https",
	)

	updates := make(chan entities.Update)

	go poller.Start(*client, updates)
	go handler.Start(*client, updates)

	select {}
}
