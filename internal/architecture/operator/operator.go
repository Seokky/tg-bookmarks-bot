// Package operator
// Manages the interaction of the poller, commander, and responser
package operator

import (
	"tg-bookmarks-bot/internal/architecture/handler"
	"tg-bookmarks-bot/internal/architecture/poller"
	"tg-bookmarks-bot/internal/domain/entities"
)

func Start() {
	updates := make(chan entities.Update)

	go poller.Start(updates)
	go handler.Start(updates)
}
