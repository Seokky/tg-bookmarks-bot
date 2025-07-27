// Package handler
// Reads updates from the updates channel and processes them
package handler

import (
	"log"
	"tg-bookmarks-bot/internal/architecture/commander"
	storerPkg "tg-bookmarks-bot/internal/architecture/storer"
	"tg-bookmarks-bot/internal/domain/entities"
)

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

func handle(storer *storerPkg.Storer, update entities.Update) {
	commander.ProcessCommand(storer, update.Message.Text, update.Message.From.Name)
	// TODO responser.Send(response)
}
