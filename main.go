// Telegram bot that storas links passed to it and allow to pick random
package main

import (
	"tg-bookmarks-bot/internal/architecture/operator"
)

func main() {
	operator.Start()

	select {}
}
