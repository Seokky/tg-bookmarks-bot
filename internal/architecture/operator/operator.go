// Package operator manages the interaction of the poller, commander, and responser
package operator

import (
	"tg-bookmarks-bot/internal/architecture/handler"
	"tg-bookmarks-bot/internal/architecture/poller"
	"tg-bookmarks-bot/internal/domain/entities"
)

// Start performs launch poller and handler and pass updates channel to both of them
func Start() {
	updates := make(chan entities.Update)

	go poller.Start(updates)
	go handler.Start(updates)
}
