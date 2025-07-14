// Package handler
// Reads updates from the updates channel and processes them
package handler

import (
	"fmt"
	"tg-bookmarks-bot/internal/domain/entities"
)

func Start(updates <-chan entities.Update) {
	fmt.Println("handler started")
}
