// Package handler reads updates from the updates channel and processes them
package handler

import (
	"log"
	"tg-bookmarks-bot/internal/architecture/commander"
	"tg-bookmarks-bot/internal/architecture/gate"
	"tg-bookmarks-bot/internal/architecture/responser"
	storerPkg "tg-bookmarks-bot/internal/architecture/storer"
	"tg-bookmarks-bot/internal/domain/entities"
)

// Start manages receiving updates from updates channel and pass it to handle func
func Start(updates <-chan entities.Update) {
	storer, err := storerPkg.New()
	if err != nil {
		log.Fatal(err)
	}

	for {
		update := <-updates

		// Wrap with goroutine to prevent blocking to read from updates channel
		go handle(storer, update)
	}

}

// Handle update received from polling Telegram server
func handle(storer *storerPkg.Storer, update entities.Update) {
	allowed := gate.Authenticate(update.Message.From.Name)
	if !allowed {
		responser.Send(update.Message.Chat.ID, "You are not authorized to use this bot.\nPlease contact @sykkoes")
		return
	}

	message := commander.ProcessCommand(storer, update.Message.Text, update.Message.From.Name)
	responser.Send(update.Message.Chat.ID, message)
}
