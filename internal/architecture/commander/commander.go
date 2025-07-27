// Package commander
package commander

import (
	"fmt"
	"log"
	"strings"
	storerPkg "tg-bookmarks-bot/internal/architecture/storer"
)

const (
	commandEmpty            = ""
	commandGetRandom        = "/random"
	commandGetRandomShortly = "/r"
	commandPop              = "/pop"
	commandCount            = "/count"
)

func ProcessCommand(storer *storerPkg.Storer, text, username string) (message string) {
	defer func() {
		if err := recover(); err != nil {
			logger := log.Default()
			logger.Println(err)
			message = "Unhandled error. Please, contact @sykkoes"
		}
	}()

	text = strings.TrimSpace(text)

	switch text {
	case commandEmpty:
		return "Incorrect command"
	case commandGetRandom, commandGetRandomShortly:
		bookmark, err := storer.Random(username)
		if err != nil {
			return "Cannot get random bookmark. Please, contact @sykkoes"
		}
		if bookmark == "" {
			return "You have no bookmark yet"
		}
		return bookmark
	case commandCount:
		count, err := storer.Count(username)
		if err != nil {
			return "Cannot get bookmark count. Please, contact @sykkoes"
		}
		return fmt.Sprintf("You have %d bookmarks", count)
	default:
		// TODO handle if bookmark already exists
		// TODO limit 150 for user
		err := storer.Append(username, text)
		if err != nil {
			return "Cannot append bookmark. Please, contact @sykkoes"
		}
		return "Bookmark appended successfully"
	}
}
